var fs = require("fs");
var browserify = require("browserify");
let moduleName = "shortid"
browserify({ ignoreMissing: true, standalone: moduleName })
  .transform(
      "babelify", {
          presets: ["babel-preset-es2015"],
          plugins:["babel-plugin-transform-remove-console"],
          compact:true
       }
  )
  .require(require.resolve(moduleName),{ entry: true })
  .bundle()
  .pipe(fs.createWriteStream(`${moduleName}.js`));