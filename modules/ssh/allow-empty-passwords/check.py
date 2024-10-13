import os

with open("/etc/ssh/sshd_config") as f:
    for line in f:
        if "PermitEmptyPasswords" in line:
            if "yes" in line:
                print("Empty passwords are allowed")
                exit(-1)
            else:
                print("Empty passwords are not allowed")
                exit(0)
            break
    else:
        print("PermitEmptyPasswords not found")
        exit(-2)
