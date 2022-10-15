### What is this?

This is a script which allows you to deploy from your terminal to your Kubernetes cluster.

Synopsis
```
kdeploy microservice [images size list]
```

Examples:
```
kdeploy data-generator
kdeploy risk-manager 50
```

### How to use it?

Run command specifying microservice name and optionally number of images to be listed.
Tool will interactively prompt you for selecting image to deploy.
After selection, it will gather all additionally needed info as full digest of image and using kubectl update image
in deployment/statefulset.

### How it works?

It gets images list and related info with help of gcloud command, updates image with kubectl.
For interactive selection go script is used.
