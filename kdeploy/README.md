# kdeploy
Deploy from the terminal.

> it's moved to its [own repo](https://github.com/Shumyk/kdeploy)

Searches for images of requested microservice in Google Container Registry,
Prompts you to interactively select an image for deployment (arrows navigation, search features),
And sets the selected image in the workload.
If microservice was not specified - it obtains possible repositories from the registry and prompts you to select it first.

kdeploy requires two configuration properties - registry and repository.
The registry is where to look for your images (e.x. us.gcr.io), and the repository is the path to your images.
Set them using:
kdeploy config set [registry|repository] [value]
Or  kdeploy config edit

Assumed that all workloads are of Deployment type. If some are StatefulSets, set them in configurations:
kdeploy config set statefulsets ms-events,ms-core

kdeploy remembers every deployment you made and allows you to redeploy previous images.
kdeploy --previous [microservice]

### How to make it run?

Run `go build` command in the directory to build the binary file (there is already a binary for OSX, you can use it).
Then place it on at `/usr/local/bin`.
