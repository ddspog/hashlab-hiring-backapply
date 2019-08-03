# Development Container
FROM ddsdok/pyserve as builder

# Add all files
COPY . .

# Exposer default port
EXPOSE 5002

# Execute binary, in Heroku as well
ENTRYPOINT [ "python", "app.py" ]