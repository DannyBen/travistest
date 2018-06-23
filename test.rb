require 'yaml'


versions = {
  'python'    => `python --version 2>&1`.strip,
  # 'python3.4' => `python --version 2>&1`.strip,
  'pip'       => `pip --version 2>&1`.strip,
  'pip3'      => `pip3 --version 2>&1`.strip,
  # 'pip3.4'    => `pip3.4 --version 2>&1`.strip,
  'ruby'      => `ruby -v 2>&1`.strip,
}

puts versions.to_yaml