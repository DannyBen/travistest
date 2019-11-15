require 'yaml'
require 'lp'

yaml = YAML.load_file '.travis.yml'
cmd = yaml['script'][1]
p system cmd
