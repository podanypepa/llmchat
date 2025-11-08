# Mistral Image Generation Example

This example demonstrates how to use the `llmchat/mistral` package to generate an image using Mistral AI's image model and save it to disk.

## How It Works

1.  **Client Initialization**: The script creates a new client to communicate with the Mistral API. An API key is required for authentication.
2.  **Request Creation**: It defines an `ImageRequest` with a model like `pixtral-12b`, a prompt for the image, and the `b64_json` response format to get the image as a base64-encoded string.
3.  **Sending and Processing**: The request is sent to the API. The script then decodes the base64 data from the response and saves it as a `mistral_generated_image.png` file in the project's root directory.

## Running the Example

Before running, ensure you have set the `MISTRAL_API_KEY` environment variable with your valid Mistral API key.

```bash
export MISTRAL_API_KEY="YOUR_API_KEY"
```

Then, you can run the example from the project's root directory using the following command:

```bash
go run ./mistral/examples/image_generator/main.go
```

After successful execution, a `mistral_generated_image.png` file will appear in the root directory.
