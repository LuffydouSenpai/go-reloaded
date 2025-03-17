<h1 align="center">Welcome to my Go-Reloaded project</h1>
<br><br>
<h2>first project in the Edu01 imperative section</h2>

## Table of contents

- [Table of contents](#table-of-contents)
- [Language](#language)
- [Code installation](#code-installation)
- [What is the code for?](#what-is-the-code-for)
- [Example](#example)
- [Authors](#authors)
- [Licence](#licence)

## Language

<h3 align="left">Language Used</h3>
<div align="left">
  <img src="https://openupthecloud.com/wp-content/uploads/2020/01/Golang.png?ezimgfmt=ng%3Awebp%2Fngcb2%2Frs%3Adevice%2Frscb2-1" height="80" width="150" alt="Go logo"  />
</div>
<br><br>

## Code installation


> Open a new folder on the desktop  <br>
> Open with *VsCode* <br>
> Open terminal <br>
> Retrieve project : `git clone https://zone01normandie.org/git/wbello/go-reloaded.git` <br>

<br>

## What is the code for?

<p>the objective of making a simple text completion/editing/auto-correction tool:</p>
- <code>(hex)</code> after a nomber for convert hexa in decimal<br>
- <code>(bin)</code> after a nomber for convert binaire in decimal <br>
- <code>(up)</code> after a word for Uppercase it <br>
- <code>(low)</code> after a word for Lowercase it <br>
- <code>(cap)</code> after a word for Capitalized it <br>
<br>

<p>For (up), (low), (cap) we can add a number :</p>
- <code>(up, 3)</code> UpperCase 3 word before <br>
- <code>(low, 3)</code> Lowercase 3 word before <br>
- <code>(cap, 3)</code> Capitalized 3 word before <br>
<br>

<p>Every instance of the punctuations <code>.</code>, <code>,</code>, <code>!</code>, <code>?</code>, <code>:</code> and <code>;</code> should be close to the previous word and with space apart from the next one.</p>
<p>Except if there are groups of punctuation like: <code>...</code> or <code>!?</code>.<p>
<br>

<p>The punctuation mark <code>'</code> will always be found with another instance of it and they should be placed to the right and left of the word in the middle of them, without any spaces.</p>
<br>

<p>If there are more than one word between the two <code>' '</code> marks, the program should place the marks next to the corresponding words</p>
<br>

<p>Every instance of <code>a</code> should be turned into <code>an</code> if the next word begins with a vowel (<code>a</code>, <code>e</code>, <code>i</code>, <code>o</code>, <code>u</code>) or a <code>h</code>.</p>

<h2>How does the code work?</h2>
<p>First and foremost:</p>
<p>Fill in the sample.txt file with your sentence as follows:</p>
<pre>
  <span>this (cap) is an example (up) of a sentence</span>
</pre>

<p>Then enter the following command in the terminal: <code>go run . "sample.txt" "result.txt" </code></p>

<br>

## Example

<p>If your data.txt contains the following:</p>
<pre>
  <span>it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.</span>
<span> Simply add 42 (hex) and 10 (bin) and you will see the result is 68.</span>
  <span>There is no greater agony than bearing a untold story inside you.</span>
  <span>Punctuation tests are ... kinda boring ,don't you think !?</span>
</pre>

<p>Your result will be:</p>
<pre>
  <span>It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair. </span>
  <span>Simply add 66 and 2 and you will see the result is 68.</span>
  <span>There is no greater agony than bearing an untold story inside you. </span>
  <span>Punctuation tests are... kinda boring, don't you think!?</span>
</pre>

## Authors

+ wbello [Gitea](https://zone01normandie.org/git/wbello) 

## Licence

<p>royalty-free</p>