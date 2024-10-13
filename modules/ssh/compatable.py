import platform

if (platform.system() != "Linux"):
    print("Only supports linux")
    exit(-1)

# check if /etc/ssh/sshd_config exists
import os
if not os.path.exists("/etc/ssh/sshd_config"):
    print("sshd_config not found")
    exit(-1)
