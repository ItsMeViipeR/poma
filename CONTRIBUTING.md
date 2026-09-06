# Contributing to Poma

Thank you for your interest in contributing to **poma**! 

To maintain code quality, maintainability, and security, please review the following guidelines before submitting a Pull Request.

---

## 🚫 No AI-Generated Code / "Vibe Coding"

To preserve the integrity and quality of this repository:

- **AI-generated contributions are strictly prohibited.** Do not submit PRs generated using ChatGPT, Copilot, Claude, or other LLMs/AI tools to avoid vibe coding.
- **Understand what you submit.** You must fully comprehend every line of code, refactor, or test you submit and be able to explain or adjust it during review.
- **Low-effort PRs will be closed immediately.** Automated, unverified, or bulk AI-generated PRs will be rejected without review.

We value human craftsmanship, clear logic, and intentional design choices over automated volume.

---

## 🛠️ How to Contribute

1. **Fork the repository** and create your branch from `main`:
```bash
git checkout -b feature/my-cool-feature
```
2. **Keep PRs focused.** One feature or bugfix per Pull Request makes review much faster.
3. **Test your changes.** Ensure all existing tests pass and add new ones if applicable:
```bash
go test ./...
```
4. **Follow standard Go idioms** (`gofmt`, clear variable naming, proper error handling).
5. **Open a Pull Request** with a clear title and description explaining what was changed and why.

Thank you for building high-quality, human-written software!
