int main()
{ int hh,mm,ss;
char type;
scanf("%d:%d:%d%c",&hh,&mm,&ss,&type);
if(type=='P'){
    if(hh!=12)
    hh+=12;
}else if(hh==12)
hh=0;

printf("%02d:%02d:%02d",hh,mm,ss);
return 0;
}





