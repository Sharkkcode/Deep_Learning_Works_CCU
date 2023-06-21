# Copyright 2021 The KServe Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
import asyncio
import logging
from typing import Dict

import numpy as np
from art.attacks.evasion.square_attack import SquareAttack
from art.attacks.evasion.shadow_attack import ShadowAttack
from art.estimators.classification import BlackBoxClassifierNeuralNetwork, PyTorchClassifier
from art.utils import load_mnist

import torch.nn as nn
import torch.nn.functional as F
import torch.optim as optim

import kserve

import nest_asyncio
nest_asyncio.apply()

AVAILABLE_ADVERSARY_TYPES = ["squareattack", "shadowattack"]

class Net(nn.Module):
    def __init__(self):
        super(Net, self).__init__()
        self.conv_1 = nn.Conv2d(in_channels=1, out_channels=4, kernel_size=5, stride=1)
        self.conv_2 = nn.Conv2d(in_channels=4, out_channels=10, kernel_size=5, stride=1)
        self.fc_1 = nn.Linear(in_features=4 * 4 * 10, out_features=100)
        self.fc_2 = nn.Linear(in_features=100, out_features=10)

    def forward(self, x):
        x = F.relu(self.conv_1(x))
        x = F.max_pool2d(x, 2, 2)
        x = F.relu(self.conv_2(x))
        x = F.max_pool2d(x, 2, 2)
        x = x.view(-1, 4 * 4 * 10)
        x = F.relu(self.fc_1(x))
        x = self.fc_2(x)
        return x

class ARTModel(kserve.Model):  # pylint:disable=c-extension-no-member
    def __init__(self, name: str, predictor_host: str, adversary_type: str,
                 nb_classes: str, max_iter: str):
        super().__init__(name)
        self.name = name
        self.predictor_host = predictor_host
        if str.lower(adversary_type) not in AVAILABLE_ADVERSARY_TYPES:
            raise Exception("Invalid adversary type: %s" % adversary_type)
        self.adversary_type = adversary_type
        self.nb_classes = int(nb_classes)
        self.max_iter = int(max_iter)
        self.ready = False
        self.count = 0

        (x_train, y_train), _, _, _ = load_mnist()
        x_train = np.transpose(x_train, (0, 3, 1, 2)).astype(np.float32)

        model = Net()

        criterion = nn.CrossEntropyLoss()
        optimizer = optim.Adam(model.parameters(), lr=0.01)

        classifier = PyTorchClassifier(
            model=model,
            clip_values=(-1, 1),
            loss=criterion,
            optimizer=optimizer,
            input_shape=(1, 28, 28),
            nb_classes=self.nb_classes,
        )

        classifier.fit(x_train, y_train, batch_size=64, nb_epochs=3)

        self.classifier = classifier

    def load(self) -> bool:
        self.ready = True
        return self.ready

    def _predict(self, x):
        n_samples = len(x)
        input_image = x.reshape((n_samples, -1))
        scoring_data = {'instances': input_image.tolist()}

        loop = asyncio.get_running_loop()
        resp = loop.run_until_complete(self.predict(scoring_data))
        prediction = np.array(resp["predictions"])
        return [1 if x == prediction else 0 for x in range(0, self.nb_classes)]

    def explain(self, payload: Dict, headers: Dict[str, str] = None) -> Dict:
        image = payload["instances"][0]
        label = payload["instances"][1]
        try:
            inputs = np.array(image)
            label = np.array(label)
            logging.info("Calling explain on image of shape %s", (inputs.shape,))
        except Exception as e:
            raise Exception(
                "Failed to initialize NumPy array from inputs: %s, %s" % (e, payload["instances"]))
        try:
            adversary_type = str.lower(self.adversary_type)
            if adversary_type == "squareattack":
                classifier = BlackBoxClassifierNeuralNetwork(self._predict, inputs.shape, self.nb_classes,
                                                             channels_first=False, clip_values=(-np.inf, np.inf))
                preds = np.argmax(classifier.predict(inputs, batch_size=1))
                attack = SquareAttack(estimator=classifier, max_iter=self.max_iter)
                x_adv = attack.generate(x=inputs, y=label)

                adv_preds = np.argmax(classifier.predict(x_adv))
                l2_error = np.linalg.norm(np.reshape(x_adv[0] - inputs, [-1]))

                return {"explanations": {"adversarial_example": x_adv.tolist(), "L2 error": l2_error.tolist(),
                                         "adversarial_prediction": adv_preds.tolist(), "prediction": preds.tolist()}}
            elif adversary_type == "shadowattack":
                inputs = np.transpose(inputs, (0, 3, 1, 2)).astype(np.float32)

                preds = np.argmax(self.classifier.predict(inputs, batch_size=1))
                attack = ShadowAttack(estimator=self.classifier)
                inputs = np.expand_dims(inputs[0], axis=0)
                x_adv = attack.generate(x=inputs)

                adv_preds = np.argmax(self.classifier.predict(x_adv))
                l2_error = np.linalg.norm(np.reshape(x_adv[0] - inputs, [-1]))

                return {"explanations": {"adversarial_example": np.transpose(x_adv, (0, 2, 3, 1)).tolist(), "L2 error": l2_error.tolist(),
                                         "adversarial_prediction": adv_preds.tolist(), "prediction": preds.tolist()}}
                
        except Exception as e:
            raise Exception("Failed to explain %s" % e)
