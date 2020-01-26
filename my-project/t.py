with open('temp.txt','r') as fp:
    l=0
    for line in fp:
        t=line.split(',')
        with open('temp1.txt','a') as fp2:
            fp2.write(t[0]+'\n')
            fp2.close()
        print(t[0])
