

document.body.onload = addElement;
function addElement() {
  const markdown = document.createElement('textarea')
  markdown.id = 'markdown'
  markdown.style = "height: 300px; width: 400px ;"

  document.body.appendChild(markdown)
  const parseBtn = document.createElement('button')
  parseBtn.id = 'parse'
  parseBtn.innerHTML = '解析markdown'
  document.body.appendChild(parseBtn)

  const previewArea = document.createElement('div')
  previewArea.id = 'preview'
  document.body.appendChild(previewArea)


  const rust = import('../pkg/index.js')

  rust.then(module => {
    const btn = document.getElementById('parse')
    const previewArea = document.getElementById('preview')

    btn.addEventListener('click', () => {
      const input = document.getElementById('markdown').value
      previewArea.innerHTML = module.parse(input)
    })
  })
}


// import("../pkg/index.js").then(module => {
//   const input = '1233'
//     previewArea.innerHTML = module.parse(input)
// }).catch(console.error);


