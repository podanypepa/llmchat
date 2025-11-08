# Grok Image Generation Example

This example demonstrates how to use the `llmchat/grok` package to generate an image using xAI's Grok image model and save it to disk.

## How It Works

1.  **Client Initialization**: The script creates a new client to communicate with the Grok API. An API key is required for authentication.
2.  **Request Creation**: It defines an `ImageRequest` with a model like `grok-2-image-1212`, a prompt for the image, and the `b64_json` response format to get the image as a base64-encoded string.
3.  **Sending and Processing**: The request is sent to the API. The script then decodes the base64 data from the response and saves it as a `grok_generated_image.png` file in the project's root directory.

## Running the Example

Before running, ensure you have set the `GROK_API_KEY` environment variable with your valid Grok API key.

```bash
export GROK_API_KEY="YOUR_API_KEY"
```

Then, you can run the example from the project's root directory using the following command:

```bash
go run ./grok/examples/image_generator/main.go
```

After successful execution, a `grok_generated_image.png` file will appear in the root directory.
