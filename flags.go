package dextk

const (
	AccPublic               uint32 = 0x0001
	AccPrivate              uint32 = 0x0002
	AccProtected            uint32 = 0x0004
	AccStatic               uint32 = 0x0008
	AccFinal                uint32 = 0x0010
	AccSynchronized         uint32 = 0x0020
	AccSuper                uint32 = 0x0020
	AccVolatile             uint32 = 0x0040
	AccBridge               uint32 = 0x0040
	AccTransient            uint32 = 0x0080
	AccVarargs              uint32 = 0x0080
	AccNative               uint32 = 0x0100
	AccInterface            uint32 = 0x0200
	AccAbstract             uint32 = 0x400
	AccStrict               uint32 = 0x800
	AccSynthetic            uint32 = 0x1000
	AccAnnotation           uint32 = 0x2000
	AccEnum                 uint32 = 0x4000
	AccConstructor          uint32 = 0x10000
	AccDeclaredSynchronized uint32 = 0x20000

	ClassFlags = AccPublic | AccFinal | AccSuper | AccInterface | AccAbstract | AccSynthetic |
		AccAnnotation | AccEnum
	InnerClassFlags = AccPublic | AccPrivate | AccProtected | AccStatic | AccFinal | AccInterface |
		AccAbstract | AccSynthetic | AccAnnotation | AccEnum
	FieldFlags = AccProtected | AccProtected | AccProtected | AccStatic | AccFinal | AccVolatile | AccTransient |
		AccSynthetic | AccEnum
	MethodFlags = AccPublic | AccPrivate | AccProtected | AccStatic | AccFinal | AccSynchronized | AccBridge |
		AccVarargs | AccNative | AccAbstract | AccStrict | AccSynthetic | AccConstructor | AccDeclaredSynchronized
)
