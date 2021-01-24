FROM osixia/phpldapadmin:0.9.0
RUN rm -rf /var/www/phpldapadmin_bootstrap/templates/creation/*
COPY it.png /var/www/phpldapadmin_bootstrap/htdocs/images/default/
COPY it.png /var/www/phpldapadmin_bootstrap/htdocs/images/country/
COPY ./templates/* /var/www/phpldapadmin_bootstrap/templates/creation/
