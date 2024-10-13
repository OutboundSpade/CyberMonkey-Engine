import os

with open("/etc/ssh/sshd_config") as f:
    # replace the line with PermitEmptyPasswords with PermitEmptyPasswords yes
    lines = f.readlines()
    f.close()
    with open("/etc/ssh/sshd_config", "w") as f:
        for line in lines:
            if "PermitEmptyPasswords" in line:
                line = "PermitEmptyPasswords yes\n"
            f.write(line)
        f.close()
    # restart sshd
    os.system("systemctl restart sshd")
    print("Empty passwords are now allowed")
