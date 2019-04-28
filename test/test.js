var nx = require('next-js-core2');
var fs = require('fs');
require('../src/next-tree-md');

test('nx.treeMd', function() {
  var rs = nx.treeMd('./test/demo');
  fs.writeFileSync('./test/output_script.md', rs);
});
