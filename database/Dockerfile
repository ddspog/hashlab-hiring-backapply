FROM philippgille/serve

# Add db migration files
ADD ./schema.up.sql /srv/
ADD ./metadata.up.yaml /srv/

# Add populate script files
ADD ./check-tables.gql /srv/
ADD ./populate.gql /srv/

# Serve all files
EXPOSE 8080 8443
ENTRYPOINT ["./serve", "-d", "/srv"]