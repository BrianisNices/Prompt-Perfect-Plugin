# Prompt Perfect Plugin

This [OpenAgents](https://openagents.com/) plugin used [PromptPerfect](https://promptperfect.jina.ai/) API to enhance prompts.

## Getting Started

1. **Clone the repository:**

    ```bash
    git clone https://github.com/BrianisNices/Prompt-Perfect-Plugin.git
    cd Prompt-Perfect-Plugin
    ```

2. **Install dependencies:**

    ```bash
    go get github.com/extism/go-pdk
    ```
3. **Set "API_KEY" on main.go file :**

    ```bash
    var API_KEY string = "YOUR_API_KEY"
    ```
3. **Compile the plugin:**

    ```bash
    tinygo build -o prompt.wasm -target wasi main.go
    ```

## Usage

```bash
extism call plugin.wasm run --input '{"prompt": "YOUR_PROMPT", "targetModel": "YOUR_MODEL"}' --wasi
```
Supported Model : 'chatgpt' | 'gpt-4' | 'stablelm-tuned-alpha-7b' | 'claude' | 'cogenerate' | 'text-davinci-003' | 'dalle' | 'sd' | 'midjourney' | 'kandinsky' | 'lexica'

## Example
```bash
$ extism call prompt.wasm run --input '{"prompt": "explain money management", "targetModel": "chatgpt"}' --wasi --allow-host 'api.promptperfect.jina.ai'

Please provide a comprehensive explanation of money management, covering topics such as budgeting, saving, investing, and debt management. Your response should aim to educate and inform the reader about the principles and strategies involved in effectively managing personal finances. Please ensure that your explanation is clear, concise, and accessible to individuals with varying levels of financial knowledge, and feel free to include practical examples or tips to illustrate your points.
```
