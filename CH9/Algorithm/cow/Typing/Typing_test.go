package Typing

import "testing"

/*
描述
牛妹在练习打字，现在按照时间顺序给出牛妹按下的键（以字符串形式给出,'<'代表回退backspace，其余字符均是牛妹打的字符，字符只包含小写字母与'<'），牛妹想知道最后在屏幕上显示的文本内容是什么。若为空则返回一个空串。
在文本内容为空的时候也可以按回退backspace（在这种情况下没有任何效果）。

示例1
输入：
"acv<"
复制
返回值：
"ac"
复制
说明：
牛妹在打完"acv"之后按了回退，所以最后是"ac"
备注：
给定一个字符串s，代表牛妹所按下的按键。
1<s.length\leq10^51<s.length≤10
5
*/
func TestTyping(t *testing.T) {
	in := "r<<ij<<<vmh<<<<<p<<b<<w<<<<<<<<<<<e<v<<yw<<<<xc<<<q<<<<<d<<<u<vm<<lux<x<<b<<<<v<<<<<w<<<<p<<z<v<v<u<<b<<<<<<<e<<<s<<a<g<<y<lus<<e<<<<<d<<ix<f<<i<rp<<<l<<<x<ud<<<<<<<<<<upc<d<<<<<<<<<<<<oio<c<t<<<<w<r<a<<<<w<g<<e<<<<b<jk<f<<j<<<<<mn<<<<<x<<rtaxf<<<m<<<<<<r<<t<<<<s<<<ea<<<<<<e<<<<<f<<w<d<<<<<<<wz<<<<<<lp<a<i<<<<<<m<<vc<<<<<l<<md<<<<ttyz<urv<r<p<z<<<i<<<<x<q<<q<<<<<<sl<<<<<rz<dl<vzsk<<pzk<<<<p<<<<e<<<<l<<<<<s<c<<<s<t<<tl<<<jk<<vz<<d<<y<<<<q<<<<<n<<<e<fw<<<b<<<<z<<<pi<<r<<<<<<<n<<i<v<<<<<y<<<<<<t<y<p<q<<<y<<<<<<h<<<<<<<<z<<<<<<asg<<<<<i<w<<p<<jt<<<<m<<t<<n<c<<<nw<<<<<<<<qg<f<<x<zvl<<<<o<<<h<<z<<t<<<<<v<b<<<<<lh<<qv<<<<<<<<<<<<z<<<v<d<r<<<s<<t<<<pw<<<<<<<<<<<<<<o<f<l<<ra<gz<<q<y<o<<pn<<<<n<gb<e<<<<abg<<e<<<<<<<<<jne<<<<<<<<<n<<tqw<<<<mx<e<<<<<j<<eori<een<<jzm<<<n<o<<c<b<z<u<<<<<<<<<a<u<<yt<<<d<y<<<<g<<<<<e<<<<p<<k<<<<j<g<<<<pn<<<<<<<<<q<<<ot<d<d<<n<<t<<<fh<t<vy<<<<k<<a<<<<s<<<<<<<<<<<<<<my<<<<a<w<<<<<<<<<p<<<z<c<<<<v<<<mmomm<<<<<<f<fcaagx"
	ret := Typing(in)
	t.Log(ret)
}
func Typing(s string) string {

	nb := []byte{}
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		if b[i] == 60 && len(nb) > 0 {
			nb = nb[:len(nb)-1]
		} else {
			nb = append(nb, b[i])
		}
	}
	nnb := []byte{}
	for _, v := range nb {
		if v != 60 {
			nnb = append(nnb, v)
		}
	}
	sb := string(nnb)
	return sb
}
func TestByte(t *testing.T) {
	s := "abcv<"
	b := []byte(s)
	t.Log(b)
}
