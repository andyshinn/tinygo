// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega3208.atdf, see http://packs.download.atmel.com/

// +build avr,atmega3208

// Device information for the ATmega3208.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega3208"
	ARCH   = "AVR8X"
	FAMILY = "AVR MEGA"
)

// Interrupts
const (
	IRQ_NMI    = 1  //
	IRQ_VLM    = 2  //
	IRQ_CNT    = 3  //
	IRQ_PIT    = 4  //
	IRQ_CCL    = 5  //
	IRQ_PORT   = 6  //
	IRQ_LUNF   = 7  //
	IRQ_OVF    = 7  //
	IRQ_HUNF   = 8  //
	IRQ_LCMP0  = 9  //
	IRQ_CMP0   = 9  //
	IRQ_CMP1   = 10 //
	IRQ_LCMP1  = 10 //
	IRQ_CMP2   = 11 //
	IRQ_LCMP2  = 11 //
	IRQ_INT    = 12 //
	IRQ_INT    = 13 //
	IRQ_TWIS   = 14 //
	IRQ_TWIM   = 15 //
	IRQ_INT    = 16 //
	IRQ_RXC    = 17 //
	IRQ_DRE    = 18 //
	IRQ_TXC    = 19 //
	IRQ_PORT   = 20 //
	IRQ_AC     = 21 //
	IRQ_RESRDY = 22 //
	IRQ_WCOMP  = 23 //
	IRQ_PORT   = 24 //
	IRQ_INT    = 25 //
	IRQ_RXC    = 26 //
	IRQ_DRE    = 27 //
	IRQ_TXC    = 28 //
	IRQ_PORT   = 29 //
	IRQ_EE     = 30 //
	IRQ_RXC    = 31 //
	IRQ_DRE    = 32 //
	IRQ_TXC    = 33 //
	IRQ_PORT   = 34 //
	IRQ_PORT   = 35 //
	IRQ_max    = 35 // Highest interrupt number on this device.
)

// Peripherals
var (
	// Analog Comparator
	AC = struct {
		DACREF   __reg
		MUXCTRLA __reg
	}{
		DACREF:   0x4, // Referance scale control
		MUXCTRLA: 0x2, // Mux Control A
	}

	// Analog to Digital Converter
	ADC = struct {
		COMMAND  __reg
		CTRLE    __reg
		MUXPOS   __reg
		RESL     __reg
		RESH     __reg
		SAMPCTRL __reg
		WINHTL   __reg
		WINHTH   __reg
		WINLTL   __reg
		WINLTH   __reg
	}{
		COMMAND:  0x8,  // Command
		CTRLE:    0x4,  // Control E
		MUXPOS:   0x6,  // Positive mux input
		RESL:     0x10, // ADC Accumulator Result
		RESH:     0x10, // ADC Accumulator Result
		SAMPCTRL: 0x5,  // Sample Control
		WINHTL:   0x14, // Window comparator high threshold
		WINHTH:   0x14, // Window comparator high threshold
		WINLTL:   0x12, // Window comparator low threshold
		WINLTH:   0x12, // Window comparator low threshold
	}

	// Bod interface
	BOD = struct {
		VLMCTRLA __reg
	}{
		VLMCTRLA: 0x8, // Voltage level monitor Control
	}

	// Configurable Custom Logic
	CCL = struct {
		INTCTRL0  __reg
		LUT0CTRLA __reg
		LUT0CTRLB __reg
		LUT0CTRLC __reg
		LUT1CTRLA __reg
		LUT1CTRLB __reg
		LUT1CTRLC __reg
		LUT2CTRLA __reg
		LUT2CTRLB __reg
		LUT2CTRLC __reg
		LUT3CTRLA __reg
		LUT3CTRLB __reg
		LUT3CTRLC __reg
		SEQCTRL0  __reg
		TRUTH0    __reg
		TRUTH1    __reg
		TRUTH2    __reg
		TRUTH3    __reg
	}{
		INTCTRL0:  0x5,  // Interrupt Control 0
		LUT0CTRLA: 0x8,  // LUT Control 0 A
		LUT0CTRLB: 0x9,  // LUT Control 0 B
		LUT0CTRLC: 0xa,  // LUT Control 0 C
		LUT1CTRLA: 0xc,  // LUT Control 1 A
		LUT1CTRLB: 0xd,  // LUT Control 1 B
		LUT1CTRLC: 0xe,  // LUT Control 1 C
		LUT2CTRLA: 0x10, // LUT Control 2 A
		LUT2CTRLB: 0x11, // LUT Control 2 B
		LUT2CTRLC: 0x12, // LUT Control 2 C
		LUT3CTRLA: 0x14, // LUT Control 3 A
		LUT3CTRLB: 0x15, // LUT Control 3 B
		LUT3CTRLC: 0x16, // LUT Control 3 C
		SEQCTRL0:  0x1,  // Sequential Control 0
		TRUTH0:    0xb,  // Truth 0
		TRUTH1:    0xf,  // Truth 1
		TRUTH2:    0x13, // Truth 2
		TRUTH3:    0x17, // Truth 3
	}

	// Clock controller
	CLKCTRL = struct {
		MCLKCTRLA    __reg
		MCLKCTRLB    __reg
		MCLKLOCK     __reg
		MCLKSTATUS   __reg
		OSC20MCALIBA __reg
		OSC20MCALIBB __reg
		OSC20MCTRLA  __reg
		OSC32KCALIB  __reg
		OSC32KCTRLA  __reg
		XOSC32KCTRLA __reg
	}{
		MCLKCTRLA:    0x0,  // MCLK Control A
		MCLKCTRLB:    0x1,  // MCLK Control B
		MCLKLOCK:     0x2,  // MCLK Lock
		MCLKSTATUS:   0x3,  // MCLK Status
		OSC20MCALIBA: 0x11, // OSC20M Calibration A
		OSC20MCALIBB: 0x12, // OSC20M Calibration B
		OSC20MCTRLA:  0x10, // OSC20M Control A
		OSC32KCALIB:  0x19, // OSC32K Calibration
		OSC32KCTRLA:  0x18, // OSC32K Control A
		XOSC32KCTRLA: 0x1c, // XOSC32K Control A
	}

	// CPU
	CPU = struct {
		CCP   __reg
		RAMPZ __reg
		SPH   __reg
		SPL   __reg
		SREG  __reg
	}{
		CCP:   0x4, // Configuration Change Protection
		RAMPZ: 0xb, // Extended Z-pointer Register
		SPH:   0xe, // Stack Pointer High
		SPL:   0xd, // Stack Pointer Low
		SREG:  0xf, // Status Register
	}

	// Interrupt Controller
	CPUINT = struct {
		LVL0PRI __reg
		LVL1VEC __reg
	}{
		LVL0PRI: 0x2, // Interrupt Level 0 Priority
		LVL1VEC: 0x3, // Interrupt Level 1 Priority Vector
	}

	// CRCSCAN
	CRCSCAN = struct {
	}{}

	// Event System
	EVSYS = struct {
		CHANNEL0     __reg
		CHANNEL1     __reg
		CHANNEL2     __reg
		CHANNEL3     __reg
		CHANNEL4     __reg
		CHANNEL5     __reg
		STROBE       __reg
		USERADC0     __reg
		USERCCLLUT0A __reg
		USERCCLLUT0B __reg
		USERCCLLUT1A __reg
		USERCCLLUT1B __reg
		USERCCLLUT2A __reg
		USERCCLLUT2B __reg
		USERCCLLUT3A __reg
		USERCCLLUT3B __reg
		USEREVOUTA   __reg
		USEREVOUTB   __reg
		USEREVOUTC   __reg
		USEREVOUTD   __reg
		USEREVOUTE   __reg
		USEREVOUTF   __reg
		USERTCA0     __reg
		USERTCB0     __reg
		USERTCB1     __reg
		USERTCB2     __reg
		USERTCB3     __reg
		USERUSART0   __reg
		USERUSART1   __reg
		USERUSART2   __reg
		USERUSART3   __reg
	}{
		CHANNEL0:     0x10, // Multiplexer Channel 0
		CHANNEL1:     0x11, // Multiplexer Channel 1
		CHANNEL2:     0x12, // Multiplexer Channel 2
		CHANNEL3:     0x13, // Multiplexer Channel 3
		CHANNEL4:     0x14, // Multiplexer Channel 4
		CHANNEL5:     0x15, // Multiplexer Channel 5
		STROBE:       0x0,  // Channel Strobe
		USERADC0:     0x28, // User ADC0
		USERCCLLUT0A: 0x20, // User CCL LUT0 Event A
		USERCCLLUT0B: 0x21, // User CCL LUT0 Event B
		USERCCLLUT1A: 0x22, // User CCL LUT1 Event A
		USERCCLLUT1B: 0x23, // User CCL LUT1 Event B
		USERCCLLUT2A: 0x24, // User CCL LUT2 Event A
		USERCCLLUT2B: 0x25, // User CCL LUT2 Event B
		USERCCLLUT3A: 0x26, // User CCL LUT3 Event A
		USERCCLLUT3B: 0x27, // User CCL LUT3 Event B
		USEREVOUTA:   0x29, // User EVOUT Port A
		USEREVOUTB:   0x2a, // User EVOUT Port B
		USEREVOUTC:   0x2b, // User EVOUT Port C
		USEREVOUTD:   0x2c, // User EVOUT Port D
		USEREVOUTE:   0x2d, // User EVOUT Port E
		USEREVOUTF:   0x2e, // User EVOUT Port F
		USERTCA0:     0x33, // User TCA0
		USERTCB0:     0x34, // User TCB0
		USERTCB1:     0x35, // User TCB1
		USERTCB2:     0x36, // User TCB2
		USERTCB3:     0x37, // User TCB3
		USERUSART0:   0x2f, // User USART0
		USERUSART1:   0x30, // User USART1
		USERUSART2:   0x31, // User USART2
		USERUSART3:   0x32, // User USART3
	}

	// Fuses
	FUSE = struct {
		APPEND  __reg
		BODCFG  __reg
		BOOTEND __reg
		OSCCFG  __reg
		SYSCFG0 __reg
		SYSCFG1 __reg
		TCD0CFG __reg
		WDTCFG  __reg
	}{
		APPEND:  0x7, // Application Code Section End
		BODCFG:  0x1, // BOD Configuration
		BOOTEND: 0x8, // Boot Section End
		OSCCFG:  0x2, // Oscillator Configuration
		SYSCFG0: 0x5, // System Configuration 0
		SYSCFG1: 0x6, // System Configuration 1
		TCD0CFG: 0x4, // TCD0 Configuration
		WDTCFG:  0x0, // Watchdog Configuration
	}

	// General Purpose IO
	GPIO = struct {
		GPIOR0 __reg
		GPIOR1 __reg
		GPIOR2 __reg
		GPIOR3 __reg
	}{
		GPIOR0: 0x0, // General Purpose IO Register 0
		GPIOR1: 0x1, // General Purpose IO Register 1
		GPIOR2: 0x2, // General Purpose IO Register 2
		GPIOR3: 0x3, // General Purpose IO Register 3
	}

	// Lockbit
	LOCKBIT = struct {
		LOCKBIT __reg
	}{
		LOCKBIT: 0x0, // Lock Bits
	}

	// BIST in the NVMCTRL module
	NVMBIST = struct {
		ADDRPAT __reg
		DATAPAT __reg
	}{
		ADDRPAT: 0x1, // Address pattern
		DATAPAT: 0x2, // Data pattern
	}

	// Non-volatile Memory Controller
	NVMCTRL = struct {
		ADDRL __reg
		ADDRH __reg
	}{
		ADDRL: 0x8, // Address
		ADDRH: 0x8, // Address
	}

	// I/O Ports
	PORT = struct {
		DIRCLR   __reg
		DIRSET   __reg
		DIRTGL   __reg
		OUTCLR   __reg
		OUTSET   __reg
		OUTTGL   __reg
		PIN0CTRL __reg
		PIN1CTRL __reg
		PIN2CTRL __reg
		PIN3CTRL __reg
		PIN4CTRL __reg
		PIN5CTRL __reg
		PIN6CTRL __reg
		PIN7CTRL __reg
		PORTCTRL __reg
	}{
		DIRCLR:   0x2,  // Data Direction Clear
		DIRSET:   0x1,  // Data Direction Set
		DIRTGL:   0x3,  // Data Direction Toggle
		OUTCLR:   0x6,  // Output Value Clear
		OUTSET:   0x5,  // Output Value Set
		OUTTGL:   0x7,  // Output Value Toggle
		PIN0CTRL: 0x10, // Pin 0 Control
		PIN1CTRL: 0x11, // Pin 1 Control
		PIN2CTRL: 0x12, // Pin 2 Control
		PIN3CTRL: 0x13, // Pin 3 Control
		PIN4CTRL: 0x14, // Pin 4 Control
		PIN5CTRL: 0x15, // Pin 5 Control
		PIN6CTRL: 0x16, // Pin 6 Control
		PIN7CTRL: 0x17, // Pin 7 Control
		PORTCTRL: 0xa,  // Port Control
	}

	// Port Multiplexer
	PORTMUX = struct {
		CCLROUTEA    __reg
		EVSYSROUTEA  __reg
		TCAROUTEA    __reg
		TCBROUTEA    __reg
		TWISPIROUTEA __reg
		USARTROUTEA  __reg
	}{
		CCLROUTEA:    0x1, // Port Multiplexer CCL
		EVSYSROUTEA:  0x0, // Port Multiplexer EVSYS
		TCAROUTEA:    0x4, // Port Multiplexer TCA
		TCBROUTEA:    0x5, // Port Multiplexer TCB
		TWISPIROUTEA: 0x3, // Port Multiplexer TWI and SPI
		USARTROUTEA:  0x2, // Port Multiplexer USART register A
	}

	// Reset controller
	RSTCTRL = struct {
		RSTFR __reg
		SWRR  __reg
	}{
		RSTFR: 0x0, // Reset Flags
		SWRR:  0x1, // Software Reset
	}

	// Real-Time Counter
	RTC = struct {
		CLKSEL      __reg
		CMPL        __reg
		CMPH        __reg
		PITCTRLA    __reg
		PITDBGCTRL  __reg
		PITINTCTRL  __reg
		PITINTFLAGS __reg
		PITSTATUS   __reg
	}{
		CLKSEL:      0x7,  // Clock Select
		CMPL:        0xc,  // Compare
		CMPH:        0xc,  // Compare
		PITCTRLA:    0x10, // PIT Control A
		PITDBGCTRL:  0x15, // PIT Debug control
		PITINTCTRL:  0x12, // PIT Interrupt Control
		PITINTFLAGS: 0x13, // PIT Interrupt Flags
		PITSTATUS:   0x11, // PIT Status
	}

	// Signature row
	SIGROW = struct {
		CHECKSUM1  __reg
		DEVICEID0  __reg
		DEVICEID1  __reg
		DEVICEID2  __reg
		OSCCAL16M0 __reg
		OSCCAL16M1 __reg
		OSCCAL20M0 __reg
		OSCCAL20M1 __reg
		OSCCAL32K  __reg
		OSC16ERR3V __reg
		OSC16ERR5V __reg
		OSC20ERR3V __reg
		OSC20ERR5V __reg
		SERNUM0    __reg
		SERNUM1    __reg
		SERNUM2    __reg
		SERNUM3    __reg
		SERNUM4    __reg
		SERNUM5    __reg
		SERNUM6    __reg
		SERNUM7    __reg
		SERNUM8    __reg
		SERNUM9    __reg
		TEMPSENSE0 __reg
		TEMPSENSE1 __reg
	}{
		CHECKSUM1:  0x2f, // CRC Checksum Byte 1
		DEVICEID0:  0x0,  // Device ID Byte 0
		DEVICEID1:  0x1,  // Device ID Byte 1
		DEVICEID2:  0x2,  // Device ID Byte 2
		OSCCAL16M0: 0x18, // Oscillator Calibration 16 MHz Byte 0
		OSCCAL16M1: 0x19, // Oscillator Calibration 16 MHz Byte 1
		OSCCAL20M0: 0x1a, // Oscillator Calibration 20 MHz Byte 0
		OSCCAL20M1: 0x1b, // Oscillator Calibration 20 MHz Byte 1
		OSCCAL32K:  0x14, // Oscillator Calibration for 32kHz ULP
		OSC16ERR3V: 0x22, // OSC16 error at 3V
		OSC16ERR5V: 0x23, // OSC16 error at 5V
		OSC20ERR3V: 0x24, // OSC20 error at 3V
		OSC20ERR5V: 0x25, // OSC20 error at 5V
		SERNUM0:    0x3,  // Serial Number Byte 0
		SERNUM1:    0x4,  // Serial Number Byte 1
		SERNUM2:    0x5,  // Serial Number Byte 2
		SERNUM3:    0x6,  // Serial Number Byte 3
		SERNUM4:    0x7,  // Serial Number Byte 4
		SERNUM5:    0x8,  // Serial Number Byte 5
		SERNUM6:    0x9,  // Serial Number Byte 6
		SERNUM7:    0xa,  // Serial Number Byte 7
		SERNUM8:    0xb,  // Serial Number Byte 8
		SERNUM9:    0xc,  // Serial Number Byte 9
		TEMPSENSE0: 0x20, // Temperature Sensor Calibration Byte 0
		TEMPSENSE1: 0x21, // Temperature Sensor Calibration Byte 1
	}

	// Sleep Controller
	SLPCTRL = struct {
	}{}

	// Serial Peripheral Interface
	SPI = struct {
	}{}

	// System Configuration Registers
	SYSCFG = struct {
		EXTBRK __reg
		OCDM   __reg
		OCDMS  __reg
		REVID  __reg
	}{
		EXTBRK: 0x2,  // External Break
		OCDM:   0x18, // OCD Message Register
		OCDMS:  0x19, // OCD Message Status
		REVID:  0x1,  // Revision ID
	}

	// 16-bit Timer/Counter Type A
	TCA = struct {
		CMP0L    __reg
		CMP0H    __reg
		CMP0BUFL __reg
		CMP0BUFH __reg
		CMP1L    __reg
		CMP1H    __reg
		CMP1BUFL __reg
		CMP1BUFH __reg
		CMP2L    __reg
		CMP2H    __reg
		CMP2BUFL __reg
		CMP2BUFH __reg
		CTRLFCLR __reg
		CTRLFSET __reg
		PERBUFL  __reg
		PERBUFH  __reg
		HCMP0    __reg
		HCMP1    __reg
		HCMP2    __reg
		HCNT     __reg
		HPER     __reg
		LCMP0    __reg
		LCMP1    __reg
		LCMP2    __reg
		LCNT     __reg
		LPER     __reg
	}{
		CMP0L:    0x28, // Compare 0
		CMP0H:    0x28, // Compare 0
		CMP0BUFL: 0x38, // Compare 0 Buffer
		CMP0BUFH: 0x38, // Compare 0 Buffer
		CMP1L:    0x2a, // Compare 1
		CMP1H:    0x2a, // Compare 1
		CMP1BUFL: 0x3a, // Compare 1 Buffer
		CMP1BUFH: 0x3a, // Compare 1 Buffer
		CMP2L:    0x2c, // Compare 2
		CMP2H:    0x2c, // Compare 2
		CMP2BUFL: 0x3c, // Compare 2 Buffer
		CMP2BUFH: 0x3c, // Compare 2 Buffer
		CTRLFCLR: 0x6,  // Control F Clear
		CTRLFSET: 0x7,  // Control F Set
		PERBUFL:  0x36, // Period Buffer
		PERBUFH:  0x36, // Period Buffer
		HCMP0:    0x29, // High Compare
		HCMP1:    0x2b, // High Compare
		HCMP2:    0x2d, // High Compare
		HCNT:     0x21, // High Count
		HPER:     0x27, // High Period
		LCMP0:    0x28, // Low Compare
		LCMP1:    0x2a, // Low Compare
		LCMP2:    0x2c, // Low Compare
		LCNT:     0x20, // Low Count
		LPER:     0x26, // Low Period
	}

	// 16-bit Timer Type B
	TCB = struct {
		CCMPL __reg
		CCMPH __reg
	}{
		CCMPL: 0xc, // Compare or Capture
		CCMPH: 0xc, // Compare or Capture
	}

	// Two-Wire Interface
	TWI = struct {
		BRIDGECTRL __reg
		MADDR      __reg
		MBAUD      __reg
		MCTRLA     __reg
		MCTRLB     __reg
		MDATA      __reg
		MSTATUS    __reg
		SADDR      __reg
		SADDRMASK  __reg
		SCTRLA     __reg
		SCTRLB     __reg
		SDATA      __reg
		SSTATUS    __reg
	}{
		BRIDGECTRL: 0x1, // Bridge Control
		MADDR:      0x7, // Master Address
		MBAUD:      0x6, // Master Baurd Rate Control
		MCTRLA:     0x3, // Master Control A
		MCTRLB:     0x4, // Master Control B
		MDATA:      0x8, // Master Data
		MSTATUS:    0x5, // Master Status
		SADDR:      0xc, // Slave Address
		SADDRMASK:  0xe, // Slave Address Mask
		SCTRLA:     0x9, // Slave Control A
		SCTRLB:     0xa, // Slave Control B
		SDATA:      0xd, // Slave Data
		SSTATUS:    0xb, // Slave Status
	}

	// Universal Synchronous and Asynchronous Receiver and Transmitter
	USART = struct {
		BAUDL    __reg
		BAUDH    __reg
		RXDATAH  __reg
		RXDATAL  __reg
		RXPLCTRL __reg
		TXDATAH  __reg
		TXDATAL  __reg
		TXPLCTRL __reg
	}{
		BAUDL:    0x8, // Baud Rate
		BAUDH:    0x8, // Baud Rate
		RXDATAH:  0x1, // Receive Data High Byte
		RXDATAL:  0x0, // Receive Data Low Byte
		RXPLCTRL: 0xe, // IRCOM Receiver Pulse Length Control
		TXDATAH:  0x3, // Transmit Data High Byte
		TXDATAL:  0x2, // Transmit Data Low Byte
		TXPLCTRL: 0xd, // IRCOM Transmitter Pulse Length Control
	}

	// User Row
	USERROW = struct {
		USERROW0  __reg
		USERROW1  __reg
		USERROW2  __reg
		USERROW3  __reg
		USERROW4  __reg
		USERROW5  __reg
		USERROW6  __reg
		USERROW7  __reg
		USERROW8  __reg
		USERROW9  __reg
		USERROW10 __reg
		USERROW11 __reg
		USERROW12 __reg
		USERROW13 __reg
		USERROW14 __reg
		USERROW15 __reg
		USERROW16 __reg
		USERROW17 __reg
		USERROW18 __reg
		USERROW19 __reg
		USERROW20 __reg
		USERROW21 __reg
		USERROW22 __reg
		USERROW23 __reg
		USERROW24 __reg
		USERROW25 __reg
		USERROW26 __reg
		USERROW27 __reg
		USERROW28 __reg
		USERROW29 __reg
		USERROW30 __reg
		USERROW31 __reg
		USERROW32 __reg
		USERROW33 __reg
		USERROW34 __reg
		USERROW35 __reg
		USERROW36 __reg
		USERROW37 __reg
		USERROW38 __reg
		USERROW39 __reg
		USERROW40 __reg
		USERROW41 __reg
		USERROW42 __reg
		USERROW43 __reg
		USERROW44 __reg
		USERROW45 __reg
		USERROW46 __reg
		USERROW47 __reg
		USERROW48 __reg
		USERROW49 __reg
		USERROW50 __reg
		USERROW51 __reg
		USERROW52 __reg
		USERROW53 __reg
		USERROW54 __reg
		USERROW55 __reg
		USERROW56 __reg
		USERROW57 __reg
		USERROW58 __reg
		USERROW59 __reg
		USERROW60 __reg
		USERROW61 __reg
		USERROW62 __reg
		USERROW63 __reg
	}{
		USERROW0:  0x0,  // User Row Byte 0
		USERROW1:  0x1,  // User Row Byte 1
		USERROW2:  0x2,  // User Row Byte 2
		USERROW3:  0x3,  // User Row Byte 3
		USERROW4:  0x4,  // User Row Byte 4
		USERROW5:  0x5,  // User Row Byte 5
		USERROW6:  0x6,  // User Row Byte 6
		USERROW7:  0x7,  // User Row Byte 7
		USERROW8:  0x8,  // User Row Byte 8
		USERROW9:  0x9,  // User Row Byte 9
		USERROW10: 0xa,  // User Row Byte 10
		USERROW11: 0xb,  // User Row Byte 11
		USERROW12: 0xc,  // User Row Byte 12
		USERROW13: 0xd,  // User Row Byte 13
		USERROW14: 0xe,  // User Row Byte 14
		USERROW15: 0xf,  // User Row Byte 15
		USERROW16: 0x10, // User Row Byte 16
		USERROW17: 0x11, // User Row Byte 17
		USERROW18: 0x12, // User Row Byte 18
		USERROW19: 0x13, // User Row Byte 19
		USERROW20: 0x14, // User Row Byte 20
		USERROW21: 0x15, // User Row Byte 21
		USERROW22: 0x16, // User Row Byte 22
		USERROW23: 0x17, // User Row Byte 23
		USERROW24: 0x18, // User Row Byte 24
		USERROW25: 0x19, // User Row Byte 25
		USERROW26: 0x1a, // User Row Byte 26
		USERROW27: 0x1b, // User Row Byte 27
		USERROW28: 0x1c, // User Row Byte 28
		USERROW29: 0x1d, // User Row Byte 29
		USERROW30: 0x1e, // User Row Byte 30
		USERROW31: 0x1f, // User Row Byte 31
		USERROW32: 0x20, // User Row Byte 32
		USERROW33: 0x21, // User Row Byte 33
		USERROW34: 0x22, // User Row Byte 34
		USERROW35: 0x23, // User Row Byte 35
		USERROW36: 0x24, // User Row Byte 36
		USERROW37: 0x25, // User Row Byte 37
		USERROW38: 0x26, // User Row Byte 38
		USERROW39: 0x27, // User Row Byte 39
		USERROW40: 0x28, // User Row Byte 40
		USERROW41: 0x29, // User Row Byte 41
		USERROW42: 0x2a, // User Row Byte 42
		USERROW43: 0x2b, // User Row Byte 43
		USERROW44: 0x2c, // User Row Byte 44
		USERROW45: 0x2d, // User Row Byte 45
		USERROW46: 0x2e, // User Row Byte 46
		USERROW47: 0x2f, // User Row Byte 47
		USERROW48: 0x30, // User Row Byte 48
		USERROW49: 0x31, // User Row Byte 49
		USERROW50: 0x32, // User Row Byte 50
		USERROW51: 0x33, // User Row Byte 51
		USERROW52: 0x34, // User Row Byte 52
		USERROW53: 0x35, // User Row Byte 53
		USERROW54: 0x36, // User Row Byte 54
		USERROW55: 0x37, // User Row Byte 55
		USERROW56: 0x38, // User Row Byte 56
		USERROW57: 0x39, // User Row Byte 57
		USERROW58: 0x3a, // User Row Byte 58
		USERROW59: 0x3b, // User Row Byte 59
		USERROW60: 0x3c, // User Row Byte 60
		USERROW61: 0x3d, // User Row Byte 61
		USERROW62: 0x3e, // User Row Byte 62
		USERROW63: 0x3f, // User Row Byte 63
	}

	// Virtual Ports
	VPORT = struct {
	}{}

	// Voltage reference
	VREF = struct {
	}{}

	// Watch-Dog Timer
	WDT = struct {
	}{}
)

// Bitfields for AC: Analog Comparator
const (
	// DACREF: Referance scale control
	DACREF_DATA = 0xff // DAC voltage reference

	// MUXCTRLA: Mux Control A
	MUXCTRLA_INVERT = 0x80 // Invert AC Output
	MUXCTRLA_MUXNEG = 0x3  // Negative Input MUX Selection
	MUXCTRLA_MUXPOS = 0x18 // Positive Input MUX Selection
)

// Bitfields for ADC: Analog to Digital Converter
const (
	// COMMAND: Command
	COMMAND_STCONV = 0x1 // Start Conversion Operation

	// CTRLE: Control E
	CTRLE_WINCM = 0x7 // Window Comparator Mode

	// MUXPOS: Positive mux input
	MUXPOS_MUXPOS = 0x1f // Analog Channel Selection Bits

	// SAMPCTRL: Sample Control
	SAMPCTRL_SAMPLEN = 0x1f // Sample lenght
)

// Bitfields for BOD: Bod interface
const (
	// VLMCTRLA: Voltage level monitor Control
	VLMCTRLA_VLMLVL = 0x3 // voltage level monitor level
)

// Bitfields for CCL: Configurable Custom Logic
const (
	// INTCTRL0: Interrupt Control 0
	INTCTRL0_INTMODE0 = 0x3  // Interrupt Mode for LUT0
	INTCTRL0_INTMODE1 = 0xc  // Interrupt Mode for LUT1
	INTCTRL0_INTMODE2 = 0x30 // Interrupt Mode for LUT2
	INTCTRL0_INTMODE3 = 0xc0 // Interrupt Mode for LUT3

	// LUT0CTRLA: LUT Control 0 A
	LUT0CTRLA_CLKSRC  = 0xe  // Clock Source Selection
	LUT0CTRLA_EDGEDET = 0x80 // Edge Detection Enable
	LUT0CTRLA_ENABLE  = 0x1  // LUT Enable
	LUT0CTRLA_FILTSEL = 0x30 // Filter Selection
	LUT0CTRLA_OUTEN   = 0x40 // Output Enable

	// LUT0CTRLB: LUT Control 0 B
	LUT0CTRLB_INSEL0 = 0xf  // LUT Input 0 Source Selection
	LUT0CTRLB_INSEL1 = 0xf0 // LUT Input 1 Source Selection

	// LUT0CTRLC: LUT Control 0 C
	LUT0CTRLC_INSEL2 = 0xf // LUT Input 2 Source Selection

	// LUT1CTRLA: LUT Control 1 A
	LUT1CTRLA_CLKSRC  = 0xe  // Clock Source Selection
	LUT1CTRLA_EDGEDET = 0x80 // Edge Detection Enable
	LUT1CTRLA_ENABLE  = 0x1  // LUT Enable
	LUT1CTRLA_FILTSEL = 0x30 // Filter Selection
	LUT1CTRLA_OUTEN   = 0x40 // Output Enable

	// LUT1CTRLB: LUT Control 1 B
	LUT1CTRLB_INSEL0 = 0xf  // LUT Input 0 Source Selection
	LUT1CTRLB_INSEL1 = 0xf0 // LUT Input 1 Source Selection

	// LUT1CTRLC: LUT Control 1 C
	LUT1CTRLC_INSEL2 = 0xf // LUT Input 2 Source Selection

	// LUT2CTRLA: LUT Control 2 A
	LUT2CTRLA_CLKSRC  = 0xe  // Clock Source Selection
	LUT2CTRLA_EDGEDET = 0x80 // Edge Detection Enable
	LUT2CTRLA_ENABLE  = 0x1  // LUT Enable
	LUT2CTRLA_FILTSEL = 0x30 // Filter Selection
	LUT2CTRLA_OUTEN   = 0x40 // Output Enable

	// LUT2CTRLB: LUT Control 2 B
	LUT2CTRLB_INSEL0 = 0xf  // LUT Input 0 Source Selection
	LUT2CTRLB_INSEL1 = 0xf0 // LUT Input 1 Source Selection

	// LUT2CTRLC: LUT Control 2 C
	LUT2CTRLC_INSEL2 = 0xf // LUT Input 2 Source Selection

	// LUT3CTRLA: LUT Control 3 A
	LUT3CTRLA_CLKSRC  = 0xe  // Clock Source Selection
	LUT3CTRLA_EDGEDET = 0x80 // Edge Detection Enable
	LUT3CTRLA_ENABLE  = 0x1  // LUT Enable
	LUT3CTRLA_FILTSEL = 0x30 // Filter Selection
	LUT3CTRLA_OUTEN   = 0x40 // Output Enable

	// LUT3CTRLB: LUT Control 3 B
	LUT3CTRLB_INSEL0 = 0xf  // LUT Input 0 Source Selection
	LUT3CTRLB_INSEL1 = 0xf0 // LUT Input 1 Source Selection

	// LUT3CTRLC: LUT Control 3 C
	LUT3CTRLC_INSEL2 = 0xf // LUT Input 2 Source Selection

	// SEQCTRL0: Sequential Control 0
	SEQCTRL0_SEQSEL = 0x7 // Sequential Selection
)

// Bitfields for CLKCTRL: Clock controller
const (
	// MCLKCTRLA: MCLK Control A
	MCLKCTRLA_CLKOUT = 0x80 // System clock out
	MCLKCTRLA_CLKSEL = 0x3  // clock select

	// MCLKCTRLB: MCLK Control B
	MCLKCTRLB_PDIV = 0x1e // Prescaler division
	MCLKCTRLB_PEN  = 0x1  // Prescaler enable

	// MCLKLOCK: MCLK Lock
	MCLKLOCK_LOCKEN = 0x1 // lock ebable

	// MCLKSTATUS: MCLK Status
	MCLKSTATUS_EXTS     = 0x80 // External Clock status
	MCLKSTATUS_OSC20MS  = 0x10 // 20MHz oscillator status
	MCLKSTATUS_OSC32KS  = 0x20 // 32KHz oscillator status
	MCLKSTATUS_SOSC     = 0x1  // System Oscillator changing
	MCLKSTATUS_XOSC32KS = 0x40 // 32.768 kHz Crystal Oscillator status

	// OSC20MCALIBA: OSC20M Calibration A
	OSC20MCALIBA_CALSEL20M = 0x80 // Calibration freq select
	OSC20MCALIBA_CAL20M    = 0x7f // Calibration

	// OSC20MCALIBB: OSC20M Calibration B
	OSC20MCALIBB_LOCK       = 0x80 // Lock
	OSC20MCALIBB_TEMPCAL20M = 0xf  // Oscillator temperature coefficient

	// OSC20MCTRLA: OSC20M Control A
	OSC20MCTRLA_RUNSTDBY = 0x2 // Run standby

	// OSC32KCALIB: OSC32K Calibration
	OSC32KCALIB_CAL32K = 0x3f // Calibration

	// OSC32KCTRLA: OSC32K Control A
	OSC32KCTRLA_RUNSTDBY = 0x2 // Run standby

	// XOSC32KCTRLA: XOSC32K Control A
	XOSC32KCTRLA_CSUT     = 0x30 // Crystal startup time
	XOSC32KCTRLA_ENABLE   = 0x1  // Enable
	XOSC32KCTRLA_RUNSTDBY = 0x2  // Run standby
	XOSC32KCTRLA_SEL      = 0x4  // Select
)

// Bitfields for CPU: CPU
const (
	// CCP: Configuration Change Protection
	CCP_CCP = 0xff // CCP signature

	// SREG: Status Register
	SREG_C = 0x1  // Carry Flag
	SREG_H = 0x20 // Half Carry Flag
	SREG_I = 0x80 // Global Interrupt Enable Flag
	SREG_N = 0x4  // Negative Flag
	SREG_S = 0x10 // N Exclusive Or V Flag
	SREG_T = 0x40 // Transfer Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_Z = 0x2  // Zero Flag
)

// Bitfields for CPUINT: Interrupt Controller
const (
	// LVL0PRI: Interrupt Level 0 Priority
	LVL0PRI_LVL0PRI = 0xff // Interrupt Level Priority

	// LVL1VEC: Interrupt Level 1 Priority Vector
	LVL1VEC_LVL1VEC = 0xff // Interrupt Vector with High Priority
)

// Bitfields for EVSYS: Event System
const (
	// CHANNEL0: Multiplexer Channel 0
	CHANNEL0_GENERATOR = 0xff // Generator selector

	// CHANNEL1: Multiplexer Channel 1
	CHANNEL1_GENERATOR = 0xff // Generator selector

	// CHANNEL2: Multiplexer Channel 2
	CHANNEL2_GENERATOR = 0xff // Generator selector

	// CHANNEL3: Multiplexer Channel 3
	CHANNEL3_GENERATOR = 0xff // Generator selector

	// CHANNEL4: Multiplexer Channel 4
	CHANNEL4_GENERATOR = 0xff // Generator selector

	// CHANNEL5: Multiplexer Channel 5
	CHANNEL5_GENERATOR = 0xff // Generator selector

	// STROBE: Channel Strobe
	STROBE_STROBE0 = 0xff // Software event on channels

	// USERADC0: User ADC0
	USERADC0_CHANNEL = 0xff // Channel selector

	// USERCCLLUT0A: User CCL LUT0 Event A
	USERCCLLUT0A_CHANNEL = 0xff // Channel selector

	// USERCCLLUT0B: User CCL LUT0 Event B
	USERCCLLUT0B_CHANNEL = 0xff // Channel selector

	// USERCCLLUT1A: User CCL LUT1 Event A
	USERCCLLUT1A_CHANNEL = 0xff // Channel selector

	// USERCCLLUT1B: User CCL LUT1 Event B
	USERCCLLUT1B_CHANNEL = 0xff // Channel selector

	// USERCCLLUT2A: User CCL LUT2 Event A
	USERCCLLUT2A_CHANNEL = 0xff // Channel selector

	// USERCCLLUT2B: User CCL LUT2 Event B
	USERCCLLUT2B_CHANNEL = 0xff // Channel selector

	// USERCCLLUT3A: User CCL LUT3 Event A
	USERCCLLUT3A_CHANNEL = 0xff // Channel selector

	// USERCCLLUT3B: User CCL LUT3 Event B
	USERCCLLUT3B_CHANNEL = 0xff // Channel selector

	// USEREVOUTA: User EVOUT Port A
	USEREVOUTA_CHANNEL = 0xff // Channel selector

	// USEREVOUTB: User EVOUT Port B
	USEREVOUTB_CHANNEL = 0xff // Channel selector

	// USEREVOUTC: User EVOUT Port C
	USEREVOUTC_CHANNEL = 0xff // Channel selector

	// USEREVOUTD: User EVOUT Port D
	USEREVOUTD_CHANNEL = 0xff // Channel selector

	// USEREVOUTE: User EVOUT Port E
	USEREVOUTE_CHANNEL = 0xff // Channel selector

	// USEREVOUTF: User EVOUT Port F
	USEREVOUTF_CHANNEL = 0xff // Channel selector

	// USERTCA0: User TCA0
	USERTCA0_CHANNEL = 0xff // Channel selector

	// USERTCB0: User TCB0
	USERTCB0_CHANNEL = 0xff // Channel selector

	// USERTCB1: User TCB1
	USERTCB1_CHANNEL = 0xff // Channel selector

	// USERTCB2: User TCB2
	USERTCB2_CHANNEL = 0xff // Channel selector

	// USERTCB3: User TCB3
	USERTCB3_CHANNEL = 0xff // Channel selector

	// USERUSART0: User USART0
	USERUSART0_CHANNEL = 0xff // Channel selector

	// USERUSART1: User USART1
	USERUSART1_CHANNEL = 0xff // Channel selector

	// USERUSART2: User USART2
	USERUSART2_CHANNEL = 0xff // Channel selector

	// USERUSART3: User USART3
	USERUSART3_CHANNEL = 0xff // Channel selector
)

// Bitfields for FUSE: Fuses
const (
	// BODCFG: BOD Configuration
	BODCFG_ACTIVE   = 0xc  // BOD Operation in Active Mode
	BODCFG_LVL      = 0xe0 // BOD Level
	BODCFG_SAMPFREQ = 0x10 // BOD Sample Frequency
	BODCFG_SLEEP    = 0x3  // BOD Operation in Sleep Mode

	// OSCCFG: Oscillator Configuration
	OSCCFG_FREQSEL = 0x3  // Frequency Select
	OSCCFG_OSCLOCK = 0x80 // Oscillator Lock

	// SYSCFG0: System Configuration 0
	SYSCFG0_CRCSRC    = 0xc0 // CRC Source
	SYSCFG0_EESAVE    = 0x1  // EEPROM Save
	SYSCFG0_RSTPINCFG = 0x8  // Reset Pin Configuration

	// SYSCFG1: System Configuration 1
	SYSCFG1_SUT = 0x7 // Startup Time

	// TCD0CFG: TCD0 Configuration
	TCD0CFG_CMPA   = 0x1  // Compare A Default Output Value
	TCD0CFG_CMPAEN = 0x10 // Compare A Output Enable
	TCD0CFG_CMPB   = 0x2  // Compare B Default Output Value
	TCD0CFG_CMPBEN = 0x20 // Compare B Output Enable
	TCD0CFG_CMPC   = 0x4  // Compare C Default Output Value
	TCD0CFG_CMPCEN = 0x40 // Compare C Output Enable
	TCD0CFG_CMPD   = 0x8  // Compare D Default Output Value
	TCD0CFG_CMPDEN = 0x80 // Compare D Output Enable

	// WDTCFG: Watchdog Configuration
	WDTCFG_PERIOD = 0xf  // Watchdog Timeout Period
	WDTCFG_WINDOW = 0xf0 // Watchdog Window Timeout Period
)

// Bitfields for LOCKBIT: Lockbit
const (
	// LOCKBIT: Lock Bits
	LOCKBIT_LB = 0xff // Lock Bits
)

// Bitfields for NVMBIST: BIST in the NVMCTRL module
const (
	// ADDRPAT: Address pattern
	ADDRPAT_AMODE = 0x70 // Address mode
	ADDRPAT_XMODE = 0x3  // X address mode
	ADDRPAT_YMODE = 0xc  // Y address mode

	// DATAPAT: Data pattern
	DATAPAT_PATTERN = 0x3 // Data check pattern
	END_END         = 0xffffff
)

// Bitfields for PORT: I/O Ports
const (
	// PIN0CTRL: Pin 0 Control
	PIN0CTRL_INVEN    = 0x80 // Inverted I/O Enable
	PIN0CTRL_ISC      = 0x7  // Input/Sense Configuration
	PIN0CTRL_PULLUPEN = 0x8  // Pullup enable

	// PIN1CTRL: Pin 1 Control
	PIN1CTRL_INVEN    = 0x80 // Inverted I/O Enable
	PIN1CTRL_ISC      = 0x7  // Input/Sense Configuration
	PIN1CTRL_PULLUPEN = 0x8  // Pullup enable

	// PIN2CTRL: Pin 2 Control
	PIN2CTRL_INVEN    = 0x80 // Inverted I/O Enable
	PIN2CTRL_ISC      = 0x7  // Input/Sense Configuration
	PIN2CTRL_PULLUPEN = 0x8  // Pullup enable

	// PIN3CTRL: Pin 3 Control
	PIN3CTRL_INVEN    = 0x80 // Inverted I/O Enable
	PIN3CTRL_ISC      = 0x7  // Input/Sense Configuration
	PIN3CTRL_PULLUPEN = 0x8  // Pullup enable

	// PIN4CTRL: Pin 4 Control
	PIN4CTRL_INVEN    = 0x80 // Inverted I/O Enable
	PIN4CTRL_ISC      = 0x7  // Input/Sense Configuration
	PIN4CTRL_PULLUPEN = 0x8  // Pullup enable

	// PIN5CTRL: Pin 5 Control
	PIN5CTRL_INVEN    = 0x80 // Inverted I/O Enable
	PIN5CTRL_ISC      = 0x7  // Input/Sense Configuration
	PIN5CTRL_PULLUPEN = 0x8  // Pullup enable

	// PIN6CTRL: Pin 6 Control
	PIN6CTRL_INVEN    = 0x80 // Inverted I/O Enable
	PIN6CTRL_ISC      = 0x7  // Input/Sense Configuration
	PIN6CTRL_PULLUPEN = 0x8  // Pullup enable

	// PIN7CTRL: Pin 7 Control
	PIN7CTRL_INVEN    = 0x80 // Inverted I/O Enable
	PIN7CTRL_ISC      = 0x7  // Input/Sense Configuration
	PIN7CTRL_PULLUPEN = 0x8  // Pullup enable

	// PORTCTRL: Port Control
	PORTCTRL_SRL = 0x1 // Slew Rate Limit Enable
)

// Bitfields for PORTMUX: Port Multiplexer
const (
	// CCLROUTEA: Port Multiplexer CCL
	CCLROUTEA_LUT0 = 0x1 // CCL LUT0
	CCLROUTEA_LUT1 = 0x2 // CCL LUT1
	CCLROUTEA_LUT2 = 0x4 // CCL LUT2
	CCLROUTEA_LUT3 = 0x8 // CCL LUT3

	// EVSYSROUTEA: Port Multiplexer EVSYS
	EVSYSROUTEA_EVOUT0 = 0x1  // Event Output 0
	EVSYSROUTEA_EVOUT1 = 0x2  // Event Output 1
	EVSYSROUTEA_EVOUT2 = 0x4  // Event Output 2
	EVSYSROUTEA_EVOUT3 = 0x8  // Event Output 3
	EVSYSROUTEA_EVOUT4 = 0x10 // Event Output 4
	EVSYSROUTEA_EVOUT5 = 0x20 // Event Output 5

	// TCAROUTEA: Port Multiplexer TCA
	TCAROUTEA_TCA0 = 0x7 // Port Multiplexer TCA0

	// TCBROUTEA: Port Multiplexer TCB
	TCBROUTEA_TCB0 = 0x1 // Port Multiplexer TCB0
	TCBROUTEA_TCB1 = 0x2 // Port Multiplexer TCB1
	TCBROUTEA_TCB2 = 0x4 // Port Multiplexer TCB2
	TCBROUTEA_TCB3 = 0x8 // Port Multiplexer TCB3

	// TWISPIROUTEA: Port Multiplexer TWI and SPI
	TWISPIROUTEA_SPI0 = 0x3  // Port Multiplexer SPI0
	TWISPIROUTEA_TWI0 = 0x30 // Port Multiplexer TWI0

	// USARTROUTEA: Port Multiplexer USART register A
	USARTROUTEA_USART0 = 0x3  // Port Multiplexer USART0
	USARTROUTEA_USART1 = 0xc  // Port Multiplexer USART1
	USARTROUTEA_USART2 = 0x30 // Port Multiplexer USART2
	USARTROUTEA_USART3 = 0xc0 // Port Multiplexer USART3
)

// Bitfields for RSTCTRL: Reset controller
const (
	// RSTFR: Reset Flags
	RSTFR_BORF   = 0x2  // Brown out detector Reset flag
	RSTFR_EXTRF  = 0x4  // External Reset flag
	RSTFR_PORF   = 0x1  // Power on Reset flag
	RSTFR_SWRF   = 0x10 // Software Reset flag
	RSTFR_UPDIRF = 0x20 // UPDI Reset flag
	RSTFR_WDRF   = 0x8  // Watch dog Reset flag

	// SWRR: Software Reset
	SWRR_SWRE = 0x1 // Software reset enable
)

// Bitfields for RTC: Real-Time Counter
const (
	// CLKSEL: Clock Select
	CLKSEL_CLKSEL = 0x3 // Clock Select

	// PITCTRLA: PIT Control A
	PITCTRLA_PERIOD = 0x78 // Period
	PITCTRLA_PITEN  = 0x1  // Enable

	// PITDBGCTRL: PIT Debug control
	PITDBGCTRL_DBGRUN = 0x1 // Run in debug

	// PITINTCTRL: PIT Interrupt Control
	PITINTCTRL_PI = 0x1 // Periodic Interrupt

	// PITINTFLAGS: PIT Interrupt Flags
	PITINTFLAGS_PI = 0x1 // Periodic Interrupt

	// PITSTATUS: PIT Status
	PITSTATUS_CTRLBUSY = 0x1 // CTRLA Synchronization Busy Flag
)

// Bitfields for SYSCFG: System Configuration Registers
const (
	// EXTBRK: External Break
	EXTBRK_ENEXTBRK = 0x1 // External break enable

	// OCDMS: OCD Message Status
	OCDMS_OCDMR = 0x1 // OCD Message Read
)

// Bitfields for TCA: 16-bit Timer/Counter Type A
const (
	// CTRLFCLR: Control F Clear
	CTRLFCLR_CMP0BV = 0x2 // Compare 0 Buffer Valid
	CTRLFCLR_CMP1BV = 0x4 // Compare 1 Buffer Valid
	CTRLFCLR_CMP2BV = 0x8 // Compare 2 Buffer Valid
	CTRLFCLR_PERBV  = 0x1 // Period Buffer Valid

	// CTRLFSET: Control F Set
	CTRLFSET_CMP0BV = 0x2 // Compare 0 Buffer Valid
	CTRLFSET_CMP1BV = 0x4 // Compare 1 Buffer Valid
	CTRLFSET_CMP2BV = 0x8 // Compare 2 Buffer Valid
	CTRLFSET_PERBV  = 0x1 // Period Buffer Valid
)

// Bitfields for TWI: Two-Wire Interface
const (
	// BRIDGECTRL: Bridge Control
	BRIDGECTRL_ENABLE  = 0x1 // Bridge Enable
	BRIDGECTRL_FMPEN   = 0x2 // FM Plus Enable
	BRIDGECTRL_SDAHOLD = 0xc // SDA Hold Time

	// MCTRLA: Master Control A
	MCTRLA_ENABLE  = 0x1  // Enable TWI Master
	MCTRLA_QCEN    = 0x10 // Quick Command Enable
	MCTRLA_RIEN    = 0x80 // Read Interrupt Enable
	MCTRLA_SMEN    = 0x2  // Smart Mode Enable
	MCTRLA_TIMEOUT = 0xc  // Inactive Bus Timeout
	MCTRLA_WIEN    = 0x40 // Write Interrupt Enable

	// MCTRLB: Master Control B
	MCTRLB_ACKACT = 0x4 // Acknowledge Action
	MCTRLB_FLUSH  = 0x8 // Flush
	MCTRLB_MCMD   = 0x3 // Command

	// MSTATUS: Master Status
	MSTATUS_ARBLOST  = 0x8  // Arbitration Lost
	MSTATUS_BUSERR   = 0x4  // Bus Error
	MSTATUS_BUSSTATE = 0x3  // Bus State
	MSTATUS_CLKHOLD  = 0x20 // Clock Hold
	MSTATUS_RIF      = 0x80 // Read Interrupt Flag
	MSTATUS_RXACK    = 0x10 // Received Acknowledge
	MSTATUS_WIF      = 0x40 // Write Interrupt Flag

	// SADDRMASK: Slave Address Mask
	SADDRMASK_ADDREN   = 0x1  // Address Enable
	SADDRMASK_ADDRMASK = 0xfe // Address Mask

	// SCTRLA: Slave Control A
	SCTRLA_APIEN  = 0x40 // Address/Stop Interrupt Enable
	SCTRLA_DIEN   = 0x80 // Data Interrupt Enable
	SCTRLA_ENABLE = 0x1  // Enable TWI Slave
	SCTRLA_PIEN   = 0x20 // Stop Interrupt Enable
	SCTRLA_PMEN   = 0x4  // Promiscuous Mode Enable
	SCTRLA_SMEN   = 0x2  // Smart Mode Enable

	// SCTRLB: Slave Control B
	SCTRLB_ACKACT = 0x4 // Acknowledge Action
	SCTRLB_SCMD   = 0x3 // Command

	// SSTATUS: Slave Status
	SSTATUS_AP      = 0x1  // Slave Address or Stop
	SSTATUS_APIF    = 0x40 // Address/Stop Interrupt Flag
	SSTATUS_BUSERR  = 0x4  // Bus Error
	SSTATUS_CLKHOLD = 0x20 // Clock Hold
	SSTATUS_COLL    = 0x8  // Collision
	SSTATUS_DIF     = 0x80 // Data Interrupt Flag
	SSTATUS_DIR     = 0x2  // Read/Write Direction
	SSTATUS_RXACK   = 0x10 // Received Acknowledge
)

// Bitfields for USART: Universal Synchronous and Asynchronous Receiver and Transmitter
const (
	// RXDATAH: Receive Data High Byte
	RXDATAH_BUFOVF = 0x40 // Buffer Overflow
	RXDATAH_DATA8  = 0x1  // Receiver Data Register
	RXDATAH_FERR   = 0x4  // Frame Error
	RXDATAH_PERR   = 0x2  // Parity Error
	RXDATAH_RXCIF  = 0x80 // Receive Complete Interrupt Flag

	// RXDATAL: Receive Data Low Byte
	RXDATAL_DATA = 0xff // RX Data

	// RXPLCTRL: IRCOM Receiver Pulse Length Control
	RXPLCTRL_RXPL = 0x7f // Receiver Pulse Lenght

	// TXDATAH: Transmit Data High Byte
	TXDATAH_DATA8 = 0x1 // Transmit Data Register (CHSIZE=9bit)

	// TXDATAL: Transmit Data Low Byte
	TXDATAL_DATA = 0xff // Transmit Data Register

	// TXPLCTRL: IRCOM Transmitter Pulse Length Control
	TXPLCTRL_TXPL = 0xff // Transmit pulse length
)
