# SnakeScape

## Introduction

Snakescape is a website that teaches novice programmers how to build bots for the MMORPG, Runescape. This repo contains all the coding examples used on the site, as well as the code used to serve the site itself. The coding examples are in Python, while the website is written using Go, HTMX, Tailwind, and Templ. I hope you find this information useful on your journey to becoming a better developer.

## Core Technologies

As mentioned above, this project depends on some awesome technologies. Let me start by giving credit where credit is due:

### Programming Languages
- [Python](https://www.python.org/)
- [Go](https://go.dev/)

### Python
- [PyAutoGui](https://pyautogui.readthedocs.io/en/latest/)
- [OpenCV](https://docs.opencv.org/4.x/d6/d00/tutorial_py_root.html)

### Go
- [Templ](https://templ.guide/)
- [Air](https://github.com/cosmtrek/air)

### CSS and Javascript
- [Htmx](https://htmx.org/)
- [Tailwindcss](https://tailwindcss.com/)

## Build Steps

This project requires a build step. The following are commands needed to build your html and css output.

### Templ HTML Generation

With templ installed and the binary somewhere on your PATH, run the following to generate your HTML components and templates (remove --watch to simply build and not hot reload)

```bash
templ generate --watch
```

### CSS File Generation

With the [Tailwind Binary](https://tailwindcss.com/blog/standalone-cli) installed and moved somewhere on your PATH, run the following to generate your CSS output for your tailwind classes (remove --watch to simply build and not hot reload)

```bash
tailwindcss -i ./static/css/input.css -o ./static/css/output.css --watch
```

### Serving with Air

With the [Air Binary](https://github.com/cosmtrek/air) installed and moved somewhere on your PATH, run the following to serve and hot reload the application:

```bash
air
```

To configure air, you can modify .air.toml in the root of the project.