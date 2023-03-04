package main

import (
	"fmt"
	"math/big"
	"time"

	"github.com/liziblockchain/liziutils/timeUtil"
	"github.com/liziblockchain/liziutils/utils"
)

func main() {
	// tryBigIntJson()
	// tryCompress()
	// trytime()

	test_validator()
}

func generalTesting() {
	fn := utils.GetAppBaseDir()
	tmp := utils.ToJSON(fn)

	fmt.Println("info:", fn, tmp)

	fmt.Println("now:", timeUtil.GetTimestampMillsecond())

	email := "abc@gmail.com"
	fmt.Println("is valid email:", utils.IsValidEmailAddr(email))
}

func tryBigIntJson() {

	type token struct {
		Time    int64
		Balance *big.Int
	}

	btc := token{
		Time:    timeUtil.GetTimestampMillsecond(),
		Balance: big.NewInt(1234),
	}

	fmt.Println("try big int:", btc)
}

func tryCompress() {
	src := []byte("lakjsdfkajl34i874y190283upaskdfj;alskdjfopq9234789asdyfklasjdflkajshdfiklq7234y89asdf长安链将区块链执行流程标准化、模块化，推进区块链技术从手工作业模式演进到自动装配生产模式，方便用户根据不同的业务需求搭建区块链系统，为技术的规模化应用提供基础。 高效易用 多语言SDK工具及浏览器插件便捷接入，一键部署工具、可视化管理台以及多种建链工具实现低门槛上手，告警与监控体系实现高效运维，打造开发、调试、部署、运维全流程易用性。 功能丰富 支持3种身份权限体系、4种数据库、5种共识算法，适配不同使用场景，支持5种合约引擎及2种P2P网络，满足各类开发者需求；提供PB级存储引擎；可支持万级节点的大规模场景。 长安链生态联盟 中国华电 基于长安链的“物资采购链”，聚焦采购过程管理、供应商信用共享、重大设备质量追溯融合应用，赋能产业发展。 中国联通lakjsdfkajl34i874y190283upaskdfj;alskdjfopq9234789asdyfklasjdflkajshdfiklq7234y89asdf长安链将区块链执行流程标准化、模块化，推进区块链技术从手工作业模式演进到自动装配生产模式，方便用户根据不同的业务需求搭建区块链系统，为技术的规模化应用提供基础。 高效易用 多语言SDK工具及浏览器插件便捷接入，一键部署工具、可视化管理台以及多种建链工具实现低门槛上手，告警与监控体系实现高效运维，打造开发、调试、部署、运维全流程易用性。 功能丰富 支持3种身份权限体系、4种数据库、5种共识算法，适配不同使用场景，支持5种合约引擎及2种P2P网络，满足各类开发者需求；提供PB级存储引擎；可支持万级节点的大规模场景。 长安链生态联盟 中国华电 基于长安链的“物资采购链”，聚焦采购过程管理、供应商信用共享、重大设备质量追溯融合应用，赋能产业发展。 中国联通lakjsdfkajl34i874y190283upaskdfj;alskdjfopq9234789asdyfklasjdflkajshdfiklq7234y89asdf长安链将区块链执行流程标准化、模块化，推进区块链技术从手工作业模式演进到自动装配生产模式，方便用户根据不同的业务需求搭建区块链系统，为技术的规模化应用提供基础。 高效易用 多语言SDK工具及浏览器插件便捷接入，一键部署工具、可视化管理台以及多种建链工具实现低门槛上手，告警与监控体系实现高效运维，打造开发、调试、部署、运维全流程易用性。 功能丰富 支持3种身份权限体系、4种数据库、5种共识算法，适配不同使用场景，支持5种合约引擎及2种P2P网络，满足各类开发者需求；提供PB级存储引擎；可支持万级节点的大规模场景。 长安链生态联盟 中国华电 基于长安链的“物资采购链”，聚焦采购过程管理、供应商信用共享、重大设备质量追溯融合应用，赋能产业发展。 中国联通lakjsdfkajl34i874y190283upaskdfj;alskdjfopq9234789asdyfklasjdflkajshdfiklq7234y89asdf长安链将区块链执行流程标准化、模块化，推进区块链技术从手工作业模式演进到自动装配生产模式，方便用户根据不同的业务需求搭建区块链系统，为技术的规模化应用提供基础。 高效易用 多语言SDK工具及浏览器插件便捷接入，一键部署工具、可视化管理台以及多种建链工具实现低门槛上手，告警与监控体系实现高效运维，打造开发、调试、部署、运维全流程易用性。 功能丰富 支持3种身份权限体系、4种数据库、5种共识算法，适配不同使用场景，支持5种合约引擎及2种P2P网络，满足各类开发者需求；提供PB级存储引擎；可支持万级节点的大规模场景。 长安链生态联盟 中国华电 基于长安链的“物资采购链”，聚焦采购过程管理、供应商信用共享、重大设备质量追溯融合应用，赋能产业发展。 中国联通在区块链技术迅速发展的这些年，数字货币被盗事件也经常发生在我们眼前。如何保证交易安全成了区块链的一大难题。为了解决这一难题，“多重签名”应运而生。它最主要的作用就是通过其技术手段来确保交易的安全性。那到底什么是“多重签名”？它又有哪些应用场景呢？今天，就跟着Starteos一起来了解一下吧~01多重签名的概念多重签名技术就是多个用户同时对一个数字资产进行签名。可以简单的理解为，一个账户多个人拥有签名权和支付权。如果一个地址只能由一个私钥签名和支付，表现形式就是1/1；而多重签名的表现形式是m/n，也就是说一共n个私钥可以给一个账户签名，而当m个地址签名时，就可以支付一笔交易。所以，m一定是小于或等于n的。")

	result := utils.Compress(src)

	fmt.Println("--------------- compress: original:", len(src), "--result:", len(result))

	decom := utils.DeCompress(result)
	fmt.Println("de-compress: original:", len(result), "---result:", len(decom))

}

func trytime() {

	curtime := time.Now().UnixMilli()
	timeUtil.CalculateStartBoundaryTimeForHalfDay(curtime)

	timeUtil.CalculateStartBoundaryTimeForHalfDay(curtime - timeUtil.MillisecondPerDay/2)

}
