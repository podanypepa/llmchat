# DeepSeek Image Generation Example

This example demonstrates how to use the `llmchat/deepseek` package to generate an image using DeepSeek's `Janus-Pro-7B` model and save it to disk.

## How It Works

1.  **Client Initialization**: The script creates a new client to communicate with the DeepSeek API. An API key is required for authentication.
2.  **Request Creation**: It defines an `ImageRequest` with the `Janus-Pro-7B` model, a prompt for the image, and the `b64_json` response format, which returns the image as a base64-encoded string.
3.  **Sending and Processing**: The request is sent to the API. The script then checks the response, decodes the base64 data, and saves it as a `deepseek_generated_image.png` file in the project's root directory.

## Running the Example

Before running, ensure you have set the `DEEPSEEK_API_KEY` environment variable with your valid DeepSeek API key.

```bash
export DEEPSEEK_API_KEY="YOUR_API_KEY"
```

Then, you can run the example from the project's root directory using the following command:

```bash
go run ./deepseek/examples/image_generator/main.go
```

After successful execution, a `deepseek_generated_image.png` file will appear in the root directory.
