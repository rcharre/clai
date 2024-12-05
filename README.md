
# Clai Command Line Interface

Clai is a command line interface that uses Mistral's language model to generate responses based on user input. This README provides an overview of the project, its dependencies, and how to use it.

## Table of Contents

- [Introduction](#introduction)
- [Dependencies](#dependencies)
- [Usage](#usage)
- [Environment Variables](#environment-variables)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Clai is a simple command line interface that allows users to interact with Mistral's language model. It can take input either as a command line argument or through a pipe. The generated response is then printed to the console.

## Dependencies

Clai relies on the following Go packages:

- [github.com/tmc/langchaingo/llms](https://github.com/tmc/langchaingo/llms)
- [github.com/tmc/langchaingo/llms/mistral](https://github.com/tmc/langchaingo/llms/mistral)

Make sure to install these dependencies before running the application.

## Usage

To use Clai, follow these steps:

1. Set the required environment variables:
   - `MISTRAL_API_KEY`: Your Mistral API key.
   - `MISTRAL_MODEL` (optional): The Mistral model to use. Defaults to "codestral-2405" if not specified.

2. Install:
   ```
   go install github.com/rcharre/clai
   ```

3. Run the compiled binary with input:
   - As a command line argument:
     ```
     clai "Your input here"
     ```
   - Through a pipe:
     ```
     echo "Your input here" | clai
     ```
   - Or both:
     ```
     echo "Your input here" | clai "what is it?"
     ```

## Environment Variables

- `MISTRAL_API_KEY`: Your Mistral API key. This is required for authentication.
- `MISTRAL_MODEL` (optional): The Mistral model to use. Defaults to "codestral-2405" if not specified.

## Contributing

Contributions are welcome! If you have any suggestions, improvements, or bug fixes, please submit a pull request. For major changes, please open an issue first to discuss what you would like to change.

