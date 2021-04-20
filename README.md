# Gomarkpdf

Gomarkpdf is a program that transforms a Markdown file to a PDF

> Warning: Images must be inside the current directory, like assets/image.png, please, dont use external image links, download and put inside the current directory.

[Github](https://github.com/edersonferreira/markpdf)

- [Install](#install)
- [Exemple](#example)
- [Usage](#usage)
- [Using Themes](#using-themes)
- [specificalling PDF size (like A4, A5,etc)](#specificalling-pdf-size)
- [Using landscape (horizontal)](#using-landscape-horizontal)
- [Switching Path](#switching-path)
- [Adding personalizated style with inline CSS in your file](#adding-personalizated-style-with-inline-css-in-your-file)
- [Append CSS files](#append-css-files)
- [Themes](#themes)
- [Creating Themes](#creating-themes)
- [List of Themes](#list-of-themes)

# Install

## Without Golang

If you don't have [Golang](https://golang.org/) installed in your machine, do these commands:

```
git clone https://github.com/edersonferreira/gomarkpdf
cd gomarkpdf
sudo cp bin/gomarkpdf /usr/bin/gomarkpdf
```

## Golang

Install Golang with your package manager, and install Gomarkpdf

To install Gomarkpdf use:

`go get -u github.com/edersonferreira/gomarkpdf`

## AUR

Gomarkpdf can be installed using AUR, with the package `gomarkpdf-edersonferreira-git`, like:

### Yay

```
yay -S gomarkpdf-edersonferreira-git
```

### Paru

```
paru gomarkpdf-edersonferreira-git
```

## Exemple

To see this README in a PDF by Gomarkpdf [click here](https://github.com/edersonferreira/markpdf/blob/master/examples/README.pdf)

# Usage

To use Gomarkpdf in a fast way, type `gomarkpdf` and your Markdown file, simple.

`gomarkpdf myExemple.md`

## Using Themes

To use a personalizated theme, append your theme file name with the `-t` or `--theme` argument, like:

`gomarkpdf myExemple.md -theme myTheme.css`

or

`gomarkpdf myExemple.md --theme myTheme.css`

If you want to use a theme in Gomarkpdf themes list, you can only write the theme name, like:

`gomarkpdf myExemple.md -t dark`

See the theme list in the bottom of this README.

An example of theme applied to Gomarkpdf is [this README in a dark theme](https://github.com/edersonferreira/gomarkpdf/blob/master/examples/dark.pdf)

## specificalling PDF size

If you want to define a specifically format, like A3, A2,etc (the default is A4), put it in the arguments with `-s` or `--size`, like:

`gomarkpdf myExemple.md -s A5`

or

`gomarkpdf myExemple.md --size A5`

An example of paper format applied to Gomarkpdf is [this README in A5 paper format](https://github.com/edersonferreira/gomarkpdf/blob/master/examples/a5.pdf)

## Using landscape (horizontal) or portrait (vertical)

If you want to use landscape format (like in slides) you can do it passing `-o landscape` or `--orientation landscape`, and to use portrait format, you can do it passing `-o portrait` or `--orientation portrait`

`gomarkpdf myExemple.md -o landscape`

or

`gomarkpdf myExemple.md --orientation portrait`

An example of Landscape format applied to Gomarkpdf is [this README in Landscape Format](https://github.com/edersonferreira/markpdf/blob/master/examples/landscape.pdf)

## Adding personalizated style with inline CSS in your file

If you need to apply a CSS change, and you don't want to create a theme to do this, you can define this using a div, with your style, like this:

```html
<div style="color:red">
Now, my text is Red!
<div>
```

You can use this to build any style with CSS, using displays, background, color, margin, or anything.

## Append CSS Files

To append CSS files, you can use the option `-c` or `-css` with the CSS file or multiple files.

```
gomarkpdf myExample.md -c MyCSSFile.css
```

or

```
gomarkpdf myExample.md -css MyCSSFile.css
```

## Themes

The default theme is white, with the h1 centralizated, justified texts,etc. See source code in `themes` directory, or the links in the footer.


## Creating Themes

If you want to create your theme, create a CSS file with a body class, defining the font-family, background color, color of text,etc. a Table style, styles for h1, h2, h3, etc. And a style for the `<code></code>` tag. Create themes for Gomarkpdf is very simple and fast, click in a theme file in section "List of Themes" to see the CSS code.

## List of Themes

[Default Theme](https://github.com/edersonferreira/gomarkpdf/blob/master/src/themes/default.css)

[Dark Theme](https://github.com/edersonferreira/gomarkpdf/blob/master/src/themes/dark.css)

[Brazilian ABNT Theme](https://github.com/edersonferreira/gomarkpdf/blob/master/src/themes/abnt.css)

[Programmer Theme](https://github.com/edersonferreira/gomarkpdf/blob/master/src/themes/programmer.css)
