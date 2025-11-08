# Gemini Image Generation Example

This example demonstrates how to use the `llmchat/gemini` package to generate an image using Google's Gemini (Imagen) model and save it to disk.

## How It Works

1.  **Client Initialization**: The script creates a new client to communicate with the Google AI API. An API key is required for authentication.
2.  **Request Creation**: It defines an `ImageRequest` with a prompt for the image.
3.  **Sending and Processing**: The request is sent to the API endpoint configured for image generation. The script then checks the response, decodes the base64 data from the first returned image, and saves it as a `gemini_generated_image.png` file in the project's root directory.

## Running the Example

Before running, ensure you have set the `GEMINI_API_KEY` environment variable with your valid Google AI API key.

```bash
export GEMINI_API_KEY="YOUR_API_KEY"
```

Then, you can run the example from the project's root directory using the following command:

```bash
go run ./gemini/examples/image_generator/main.go
```

After successful execution, a `gemini_generated_image.png` file will appear in the root directory.
