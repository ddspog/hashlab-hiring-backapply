import os


envor = lambda a, b: os.environ[a] if a in os.environ and os.environ[a] != '' else b
