# Taste go-ora

Docker is required to taste

```sh
# docker exec -it oracle11g sqlplus -s sys/oracle as sysdba
docker exec -it oracle11g sqlplus -s system/oracle

SELECT instance_name, version, status, instance_role, logins FROM v$instance;

# service name
SELECT name FROM v$database;

# sid
SELECT instance FROM v$thread;
```

* https://blogs.oracle.com/developers/post/connecting-a-go-application-to-oracle-database
* https://github.com/sijms/go-ora
* https://github.com/jaspeen/docker-oracle-xe-11g
