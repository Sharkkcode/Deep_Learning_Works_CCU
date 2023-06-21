# Our Work
## Expectation
Add more adversarial attack option to the art explainer, we started with the ShadowAttack.

## Added / Modified Files
`docs/samples/explanation/art/mnist/artserver.yaml`<br>
Added for local testing with our own docker image, might not be necessary to commit.

`hack/quick_install_requirements.sh`<br>
Added for quick development requirements installation, for example, you can run this after created a minikube cluster to quick install all the requirements for deploying kserve locally. No need to commit this file.

`python/artexplainer.Dockerfile`<br>
Modified because of some lack of modules during docker build. No need to commit this file unless you find some reason to prove that this situation is happening to other kserve users.

`python/artexplainer/artserver/model.py`<br>
Where most of our work is located.<br>
Details of what we modified can be viewed in git history or our report and ppt, our longer expectation is to make the choosing process of attack method a modular design that can automatically suit every attack method, but we only manual added one method (ShadowAttack) for now.<br>
The Net model is for temporary testing, should be removed and replaced by the model fetched from the user. The model can either be specified by user in explainer field in yaml file, or automatically imported from predictor, which might be harder in our opinion.<br>

Good luck! 😀