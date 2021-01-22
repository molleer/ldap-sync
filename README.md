# LDAP Sync

Syncing legacy [LDAP](https://ldap.com/) user database with the newer user service [Gamma](https://gamma.chalmers.it)

## Why?

The [IT student division](https://chalmers.it/) at Chalmers University of Technology previously used the user service [Beta Account](https://github.com/cthit/chalmersit-account-rails) which was written in [Ruby on Rails](https://rubyonrails.org/) and used LDAP as a user database. This service has been replaced by Gamma which is written in [Spring Boot](https://spring.io/projects/spring-boot) and [ReactJs](https://reactjs.org/) and uses a [PostgreSQL](https://www.postgresql.org/) database. However, there are still several services which is relying directly on LDAP to work which means the data in both Gamma and the LDAP database must be synced.

<!--## How?>
