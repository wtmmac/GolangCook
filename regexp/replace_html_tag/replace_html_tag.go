package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := `<p>　<script>alert()</script>　还只是在凌晨4点，家住徐汇的王女士在手指关节的阵阵僵痛中醒来，这种每每在凌晨出现的疼痛已经折磨她将近半年了。之前她一直认为是普通的关节痛，直到最近才在正规医院被诊断为<a href="http://jbk.39.net/lfsxgjy/" target="_blank" keycmd="bindJbkUi">类风湿</a><a href="http://jbk.39.net/fsxgjy/" target="_blank" keycmd="bindJbkUi">关节炎</a>。在我国，出现同样状况把<a href="http://jbk.39.net/lfsxgjy/" target="_blank" keycmd="bindJbkUi">类风湿性关节炎</a>错当成普通关节痛的却绝非王女士一人。</p> <p>　<strong>　关节疼痛要当心，以免延误病情</strong></p><p>　　据统计，目前我国大陆地区<a href="http://zzk.39.net/zz/quanshen/a748c.html" target="_blank" keycmd="null">类风湿关节炎</a>患病率为0.2%~0.36%，患者高达500-1000万名。中国中医科学院科技合作中心中经堂专家表示，现在八成以上的类风湿性关节炎的患者通过吃药和调理可以达到病情完全缓解，或控制病情的效果，但遗憾的是，“很多人一出现关节痛，首先想到的是去看骨科，或者自己买些止痛药吃，极少人会将此类现象与风湿性疾病联系在一起。最终耽误了病情，导致关节变形，甚至残废。”</p>`

	str = strings.TrimSpace(str) //去空格
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	str = re.ReplaceAllStringFunc(str, strings.ToLower)
	//替换掉注释和一些标签
	reg := regexp.MustCompile(`<!--[^>]+>|<iframe[\S\s]+?</iframe>|<a[^>]+>|</a>|<script[\S\s]+?</script>|<div class="hzh_botleft">[\S\s]+?</div>`)
	str = reg.ReplaceAllString(str, "")
	fmt.Println(str)
}
