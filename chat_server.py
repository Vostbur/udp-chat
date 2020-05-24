import socket


clients = set()

with socket.socket(socket.AF_INET, socket.SOCK_DGRAM) as s:
    s.bind(('', 10000))
    print('Start server... Post "STOP" for end')
    
    while True:
        data, addr = s.recvfrom(1024)
        print(addr[0], addr[1])

        clients.add(addr)
        for client in clients-{addr,}:
            s.sendto(data, client)

        if data.decode().endswith('STOP'):
            print("Stop chat.")
            break
