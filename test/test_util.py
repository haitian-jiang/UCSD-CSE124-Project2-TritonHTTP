from socket import socket
import time

s = socket()
s.connect(("localhost", 8080))

def send_msg(send, recv=4096, fulltxt=False):
    msg = b'a' * send
    s.sendall(msg)
    resp = s.recv(recv)
    return resp if fulltxt else len(resp)
