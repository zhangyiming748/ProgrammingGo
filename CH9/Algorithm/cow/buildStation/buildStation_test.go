package buildStation

/*
描述
X轴上10^8个点，从左到右依次编号为0~10^8-1，相邻点距离为1，其中有n个点有特殊意义，从小到大依次为a0~an-1，其中保证a0=0.
现在需要建设收集站，有两个要求必须被满足：
1、每个有意义的点必须修建收集站。
2、相邻收集站的距离必须为1或为某个质数。
现给出n和a数组，求需要建设收集站的最小数量。
示例1
输入：
3,[0,7,11]
复制
返回值：
4
复制
说明：
在0,7,8,11处建造收集站，差值分别为7,1,3，符合要求
备注：
输入数据包含一个数n和一个长度为n的数组a，数组a的下标为0~n-1，保证a数组递增，且a0=0。
输出一个整数，表示答案。
bool iff(int n)
{
if(n==1)return 1;
if(n==2)return 1;
for(int i=2;i<=sqrt(n);i++)
{
if(n%i==0)return 0;
}
return 1;
}
int work(int n, int* a, int aLen) {
// write code here
int num=0;
for(int i=1;i<aLen;i++)
{
int len=a[i]-a[i-1];
if(iff(len))num++;
else if(len%2==0)num+=2;
else if(len%2==1)
{
if(iff(len-2))num+=2;
else num+=3;
}
}
num+=1;
int sum=num;
return sum;
}
};
*/
func work(n int, a []int) int {
	ans:=n
	for i:=0;i<n;i++{
		val:=a[i]-a[i-1]
		if isprime(val){
			continue
		}else if isprime(val-2){
			ans+=1
		}else if !(val&1)
	}
}
func isprime(v int) bool {
	for i:=0;i*i<=n;i++{

	}
}
