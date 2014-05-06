initReveal = ->
  Reveal.initialize
    controls: true
    progress: true
    history: true
    center: true
    hideAddressBar: true

    transition: 'none'
    transitionSpeed: 'fast'
    backgroundTransition: 'slide'

    width: 1024
    height: 768
    margin: 0.0
    minScale: 0.2
    maxScale: 4.0

  hljs.initHighlightingOnLoad()

  if window.location.search.match /print-pdf/gi
    link = document.createElement 'link'
    link.rel = 'stylesheet'
    link.type = 'text/css'
    link.href = 'reveal/pdf.css'
    document.getElementsByTagName('head')[0].appendChild link

main = ->
  initReveal()

main()
