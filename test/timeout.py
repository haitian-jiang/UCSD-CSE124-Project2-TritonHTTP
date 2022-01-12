from test_util import *  # got s

def timeout():
    s = get_socket()
    print(send_msg(GDRQ, s))
    print(send_msg(GDRQ, s))
    sleep(7)

def normal_close():
    s = get_socket()
    print(send_msg(GDRQ, s))
    sleep(1)
    print(send_msg(CLOSE, s))

if __name__ == "__main__":
    timeout()
    normal_close()
