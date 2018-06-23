require 'yaml'


versions = {
  'python'     => `python --version 2>&1`.strip,
  'pip'        => `pip --version 2>&1`.strip,
  'ruby'       => `ruby -v 2>&1`.strip,
  'codenamize' => `codenamize --version`.strip
}

puts versions.to_yaml