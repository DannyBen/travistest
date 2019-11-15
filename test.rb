require 'yaml'
require 'lp'

yaml = YAML.load_file '.travis.yml'
lp yaml