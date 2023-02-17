# kron


**kron** will become a tool for managing cron jobs for remote hosts. It will enable sysadmins to get/create/delete/edit cron jobs on specified list of target hosts using ssh protocol, so no dependencies will be required on the target hosts.

## manual
| command | description |
| - | - |
| **kron** hosts | shows loaded hosts | 
| **kron** get all | retrieve cron job information from all hosts |
| **kron** get HOST... | retrieve cron job information from the specified hosts |
| **kron** create all CRON_JOB | create cron job on all hosts |
| **kron** create HOST... | create cron job on specified hosts |
| **kron** delete all CRON_JOB | delete cron job from all hosts |
| **kron** delete HOST... CRON_JOB | delete cron job from specified hosts |
| **kron** edit 'OLD_CRON_JOB' 'NEW_CRON_JOB' all | edit cron job in all hosts |
| **kron** edit 'OLD_CRON_JOB' 'NEW_CRON_JOB' HOSTS... | edit cron job in specified hosts | 

## hosts file
it is a yaml formatted file resides in `/etc/kron/hosts.yaml`
### format
```yaml
- host: [ an identifier string for the host ]
  address: [ domain name or ip address of the host ]
  user: [ username of the user with sudo privileges ]
  password: [ password for that user ]
```

### example hosts.yaml file
```yaml
- host: ES-server
  address: 142.250.187.174
  user: sundar
  password: Q0VNIEJBQkEgMTIzNDUK

- host: Jenkins
  address: jenkins.devops
  user: ubuntu
  password: J3Nk!nS!

- host: k8s_bastion
  address: 172.16.1.5
  user: cem
  password: 1337
```

# contact
I aim to learn go and cli development while developing this.

any contribution is welcome! you can reach me from **cemanaral@hotmail.com**
