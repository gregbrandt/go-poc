const webpack = require('webpack')
const debug = require('debug')('app:build:webpack-watch')
const config = require('../config')

function webpackWatcher (webpackConfig, statsFormat) {
  statsFormat = statsFormat || config.compiler_stats

  return new Promise((resolve, reject) => {
    const compiler = webpack(webpackConfig)
    
    compiler.watch({ // watch options:
      aggregateTimeout: 1000, // wait so long for more changes
      poll: true, // use polling instead of native watchers
      ignore: /node_modules/
      // pass a number to set the polling interval
    },(err, stats) => {
      if (err) {
        debug('Webpack compiler encountered a fatal error.', err)
        return reject(err)
      }

      const jsonStats = stats.toJson()
      debug('Webpack compile completed.')
      debug(stats.toString(statsFormat))

      if (jsonStats.errors.length > 0) {
        debug('Webpack compiler encountered errors.')
        debug(jsonStats.errors.join('\n'))
        return reject(new Error('Webpack compiler encountered errors'))
      } else if (jsonStats.warnings.length > 0) {
        debug('Webpack compiler encountered warnings.')
        debug(jsonStats.warnings.join('\n'))
      } else {
        debug('No errors or warnings encountered.')
      }
      resolve(jsonStats)
    })
  })
}

module.exports = webpackWatcher
