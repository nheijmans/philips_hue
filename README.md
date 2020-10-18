# Go Hue Control!
The hueControl script allows you to retrieve the information from your Hue devices programmatically with the standard builtin API.
This [blog post](https://threat-hunting.ninja/posts/philips-hue-automation) accompanies this repo.

## Usage
First pull the Docker image
```
docker pull statixs/hue-control:latest
```

Retrieve the raw data JSON from all your devices to see what's available
```
docker run --rm statixs/hue-control -ip <YOUR_HUE_IPADDR> -key <YOURAPIKEY> -raw | python -m 'json.tool' # last part is to pretty print the JSON
```

Retrieve information from the lightbulb with ID 4
```
 docker run --rm statixs/hue-control -ip <YOUR_HUE_IPADDR> -key <YOURAPIKEY> -lid 4
```

Retrieve information from the temperature sensor with ID 5
```
 docker run --rm statixs/hue-control -ip <YOUR_HUE_IPADDR> -key <YOURAPIKEY> -sid 5
```

# Contributions
If you want to extend the list of features, feel free to submit a pull request! The [Docker Hub](https://hub.docker.com/r/statixs/hue-control) image is updated once merged to master automatically. 
