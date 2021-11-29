// 使用加密算法，与后端请求
import JsRSA from "jsencrypt"

// 申明js加密对象
let RSACrypto =new JsRSA()
let publicKey ="-----BEGIN RSA Public Key-----MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtquH/30p4YWPPeDtPFg520RobSLhfe2elX+UTzYiu85aFC89Wta3qmDrR8f1xzwuJvw0AHMYrTJlIxH7THYi/Plor2miF9FYF+owComwUsxy/ws8VXSDhrWmCvNcwdux55iSRrVRyQAtsBxyaJUV5/DW7qlPAhG176TEChzdtWCzz0JWa9nwr3uwyqx+9tJAV7ay85Cw8e3kfwgBYJb6yHsAUPt2uUARV/gGOB5fH3NMGyzUdYXRSRapo4iW5tLSes/gfTBKjzpbeAy5S5u0Ns+XXF+0oo2H6Bjtlbj2uTctxJfiIqD4TUw3WERHnQ6QGwaT01zP97p/bv5d+8pRGwIDAQAB-----END RSA Public Key-----"
RSACrypto.setPublicKey(publicKey)


module.exports={
	RSATest:function(){
		let SrcText=RSACrypto.encrypt("{time:3543645654654,msg:520}")
		console.log(SrcText)
		return SrcText
	}
}




