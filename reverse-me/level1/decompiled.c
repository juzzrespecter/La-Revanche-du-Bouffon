
undefined4 main(void)

{
  int iVar1;
  char local_7e [4];
  char local_7a [4];
  char local_76 [4];
  char local_72 [2];
  char local_70 [100];
  undefined4 local_c;
  
  local_c = 0;
  local_7e[0] = '_';
  local_7e[1] = '_';
  local_7e[2] = 's';
  local_7e[3] = 't';
  local_7a[0] = 'a';
  local_7a[1] = 'c';
  local_7a[2] = 'k';
  local_7a[3] = '_';
  local_76[0] = 'c';
  local_76[1] = 'h';
  local_76[2] = 'e';
  local_76[3] = 'c';
  local_72[0] = 'k';
  local_72[1] = '\0';
  printf("Please enter key: ");
  __isoc99_scanf(&DAT_00012029,local_70);
  iVar1 = strcmp(local_70,local_7e);
  if (iVar1 == 0) {
    printf("Good job.\n");
  }
  else {
    printf("Nope.\n");
  }
  return 0;
}

