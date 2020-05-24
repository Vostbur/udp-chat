import socket
import threading

flag = True

def chat(sock):
    global flag
    while flag:
        data = sock.recv(1024)
        print(data.decode())
        if data.decode().endswith('STOP'):
            flag = False


server = ('localhost', 10000)
alias = input('Input user alias: ')

with socket.socket(socket.AF_INET, socket.SOCK_DGRAM) as s:
    s.bind(('', 0))
    s.sendto('[{}] connect to chat'.format(alias).encode(), server)
    threading.Thread(target=chat, args=(s,), daemon=True).start()
    
    while flag:
        post = input()
        s.sendto('[{}] {}'.format(alias, post).encode(), server)
        if post == 'STOP':
            flag = False
