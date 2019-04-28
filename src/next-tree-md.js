(function() {
  var global = global || this || window || Function('return this')();
  var nx = global.nx || require('next-js-core2');
  var dirTree = require('directory-tree');
  var DEFAULT_OPTIONS = { extensions: /\.md/ };

  // import packages:
  require('next-repeat');

  nx.treeMd = function(inPath, inOptions) {
    var options = nx.mix(DEFAULT_OPTIONS, inOptions);
    var root = dirTree(inPath, options);
    var lines = ['## menu'];
    var getLines = function(inSpace, node) {
      var items = node.children;
      lines.push(inSpace + '+ **' + node.name + '**');
      items.forEach(function(item) {
        var children = item.children;
        var spaces = inSpace + '    ';
        if (children && children.length) {
          getLines(spaces, item);
        } else {
          lines.push(spaces + '- [' + item.name + '](' + item.path + ')');
        }
      });
    };
    getLines('', root);
    return lines.join('\n');
  };

  if (typeof module !== 'undefined' && module.exports) {
    module.exports = nx.treeMd;
  }
})();
