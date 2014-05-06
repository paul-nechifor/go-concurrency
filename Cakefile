require('coffee-script').register()
Build = require('web-build-tools').Build
{sh, cmd} = Build
jade = require 'jade'
fs = require 'fs'

config =
  debug: false

b = new Build task, config, (->),
  clean: (cb) ->
    sh 'rm -fr build; mkdir build', cb

  copyReq: (cb) ->
    sh """
      mkdir build/reveal build/js
      cd bower_components/reveal.js
      cp js/reveal.min.js css/reveal.min.css lib/js/* ../../build/reveal
      cp css/print/pdf.css ../../build/reveal
      cp plugin/highlight/highlight.js ../../build/js
    """, cb

  stylus: (cb) ->
    Build.stylus 'build/style.css', 'style.styl',
      config, cb

  jade: (cb) ->
    jade.renderFile 'index.jade', {}, (err, html) ->
      fs.writeFileSync 'build/index.html', html
      cb?()

  script: (cb) ->
    Build.browserify 'build/script.js', './script.coffee', config, cb

  build: (cb) ->
    b.run ['clean', 'copyReq', 'stylus', 'jade', 'script']
    cb?()

b.makePublic
  build: 'Build the presentation.'
