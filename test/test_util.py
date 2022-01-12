from socket import socket
from time import sleep

# good request if add b'\r\n'
gdrq = b'GET / HTTP/1.1\r\nHost: localhost\r\n'
GDRQ = gdrq + b'\r\n'
# close request
CLOSE = b'GET / HTTP/1.1\r\nHost: localhost\r\nconnection:close\r\n\r\n'

def get_socket():
    s = socket()
    s.connect(("localhost", 8080))
    return s

def send_msg(msg, s=None, recv=4096, fulltxt=True):
    keep = 1
    if not s:
        keep = 0
        s = get_socket()
    s.sendall(msg)
    resp = s.recv(recv)
    if not keep: s.close()
    return resp if fulltxt else len(resp)
