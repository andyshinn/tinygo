// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega32C1.atdf, see http://packs.download.atmel.com/

// +build avr,atmega32c1

// Device information for the ATmega32C1.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega32C1"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin, Power-on Reset, Brown-out Reset, Watchdog Reset and JTAG AVR Reset
	IRQ_ANACOMP0     = 1  // Analog Comparator 0
	IRQ_ANACOMP1     = 2  // Analog Comparator 1
	IRQ_ANACOMP2     = 3  // Analog Comparator 2
	IRQ_ANACOMP3     = 4  // Analog Comparator 3
	IRQ_PSC_FAULT    = 5  // PSC Fault
	IRQ_PSC_EC       = 6  // PSC End of Cycle
	IRQ_INT0         = 7  // External Interrupt Request 0
	IRQ_INT1         = 8  // External Interrupt Request 1
	IRQ_INT2         = 9  // External Interrupt Request 2
	IRQ_INT3         = 10 // External Interrupt Request 3
	IRQ_TIMER1_CAPT  = 11 // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 12 // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 13 // Timer/Counter1 Compare Match B
	IRQ_TIMER1_OVF   = 14 // Timer1/Counter1 Overflow
	IRQ_TIMER0_COMPA = 15 // Timer/Counter0 Compare Match A
	IRQ_TIMER0_COMPB = 16 // Timer/Counter0 Compare Match B
	IRQ_TIMER0_OVF   = 17 // Timer/Counter0 Overflow
	IRQ_CAN_INT      = 18 // CAN MOB, Burst, General Errors
	IRQ_CAN_TOVF     = 19 // CAN Timer Overflow
	IRQ_LIN_TC       = 20 // LIN Transfer Complete
	IRQ_LIN_ERR      = 21 // LIN Error
	IRQ_PCINT0       = 22 // Pin Change Interrupt Request 0
	IRQ_PCINT1       = 23 // Pin Change Interrupt Request 1
	IRQ_PCINT2       = 24 // Pin Change Interrupt Request 2
	IRQ_PCINT3       = 25 // Pin Change Interrupt Request 3
	IRQ_SPI_STC      = 26 // SPI Serial Transfer Complete
	IRQ_ADC          = 27 // ADC Conversion Complete
	IRQ_WDT          = 28 // Watchdog Time-Out Interrupt
	IRQ_EE_READY     = 29 // EEPROM Ready
	IRQ_SPM_READY    = 30 // Store Program Memory Read
	IRQ_max          = 30 // Highest interrupt number on this device.
)

// Peripherals
var (
	// Fuses
	FUSE = struct {
		EXTENDED __reg
		HIGH     __reg
		LOW      __reg
	}{
		EXTENDED: 0x2,
		HIGH:     0x1,
		LOW:      0x0,
	}

	// Lockbits
	LOCKBIT = struct {
		LOCKBIT __reg
	}{
		LOCKBIT: 0x0,
	}

	// I/O Port
	PORT = struct {
		PORTB __reg
		DDRB  __reg
		PINB  __reg
		PORTC __reg
		DDRC  __reg
		PINC  __reg
		PORTD __reg
		DDRD  __reg
		PIND  __reg
		PORTE __reg
		DDRE  __reg
		PINE  __reg
	}{
		PORTB: 0x25, // Port B Data Register
		DDRB:  0x24, // Port B Data Direction Register
		PINB:  0x23, // Port B Input Pins
		PORTC: 0x28, // Port C Data Register
		DDRC:  0x27, // Port C Data Direction Register
		PINC:  0x26, // Port C Input Pins
		PORTD: 0x2b, // Port D Data Register
		DDRD:  0x2a, // Port D Data Direction Register
		PIND:  0x29, // Port D Input Pins
		PORTE: 0x2e, // Port E Data Register
		DDRE:  0x2d, // Port E Data Direction Register
		PINE:  0x2c, // Port E Input Pins
	}

	// Controller Area Network
	CAN = struct {
		CANGCON  __reg
		CANGSTA  __reg
		CANGIT   __reg
		CANGIE   __reg
		CANEN2   __reg
		CANEN1   __reg
		CANIE2   __reg
		CANIE1   __reg
		CANSIT2  __reg
		CANSIT1  __reg
		CANBT1   __reg
		CANBT2   __reg
		CANBT3   __reg
		CANTCON  __reg
		CANTIML  __reg
		CANTIMH  __reg
		CANTTCL  __reg
		CANTTCH  __reg
		CANTEC   __reg
		CANREC   __reg
		CANHPMOB __reg
		CANPAGE  __reg
		CANSTMOB __reg
		CANCDMOB __reg
		CANIDT4  __reg
		CANIDT3  __reg
		CANIDT2  __reg
		CANIDT1  __reg
		CANIDM4  __reg
		CANIDM3  __reg
		CANIDM2  __reg
		CANIDM1  __reg
		CANSTML  __reg
		CANSTMH  __reg
		CANMSG   __reg
	}{
		CANGCON:  0xd8, // CAN General Control Register
		CANGSTA:  0xd9, // CAN General Status Register
		CANGIT:   0xda, // CAN General Interrupt Register Flags
		CANGIE:   0xdb, // CAN General Interrupt Enable Register
		CANEN2:   0xdc, // Enable MOb Register 2
		CANEN1:   0xdd, // Enable MOb Register 1(empty)
		CANIE2:   0xde, // Enable Interrupt MOb Register 2
		CANIE1:   0xdf, // Enable Interrupt MOb Register 1 (empty)
		CANSIT2:  0xe0, // CAN Status Interrupt MOb Register 2
		CANSIT1:  0xe1, // CAN Status Interrupt MOb Register 1 (empty)
		CANBT1:   0xe2, // CAN Bit Timing Register 1
		CANBT2:   0xe3, // CAN Bit Timing Register 2
		CANBT3:   0xe4, // CAN Bit Timing Register 3
		CANTCON:  0xe5, // Timer Control Register
		CANTIML:  0xe6, // Timer Register
		CANTIMH:  0xe6, // Timer Register
		CANTTCL:  0xe8, // TTC Timer Register
		CANTTCH:  0xe8, // TTC Timer Register
		CANTEC:   0xea, // Transmit Error Counter Register
		CANREC:   0xeb, // Receive Error Counter Register
		CANHPMOB: 0xec, // Highest Priority MOb Register
		CANPAGE:  0xed, // Page MOb Register
		CANSTMOB: 0xee, // MOb Status Register
		CANCDMOB: 0xef, // MOb Control and DLC Register
		CANIDT4:  0xf0, // Identifier Tag Register 4
		CANIDT3:  0xf1, // Identifier Tag Register 3
		CANIDT2:  0xf2, // Identifier Tag Register 2
		CANIDT1:  0xf3, // Identifier Tag Register 1
		CANIDM4:  0xf4, // Identifier Mask Register 4
		CANIDM3:  0xf5, // Identifier Mask Register 3
		CANIDM2:  0xf6, // Identifier Mask Register 2
		CANIDM1:  0xf7, // Identifier Mask Register 1
		CANSTML:  0xf8, // Time Stamp Register
		CANSTMH:  0xf8, // Time Stamp Register
		CANMSG:   0xfa, // Message Data Register
	}

	// Analog Comparator
	AC = struct {
		AC0CON __reg
		AC1CON __reg
		AC2CON __reg
		AC3CON __reg
		ACSR   __reg
	}{
		AC0CON: 0x94, // Analog Comparator 0 Control Register
		AC1CON: 0x95, // Analog Comparator 1 Control Register
		AC2CON: 0x96, // Analog Comparator 2 Control Register
		AC3CON: 0x97, // Analog Comparator 3 Control Register
		ACSR:   0x50, // Analog Comparator Status Register
	}

	// Digital-to-Analog Converter
	DAC = struct {
		DACL  __reg
		DACH  __reg
		DACON __reg
	}{
		DACL:  0x91, // DAC Data Register
		DACH:  0x91, // DAC Data Register
		DACON: 0x90, // DAC Control Register
	}

	// CPU Registers
	CPU = struct {
		SPMCSR __reg
		SREG   __reg
		SPL    __reg
		SPH    __reg
		MCUCR  __reg
		MCUSR  __reg
		OSCCAL __reg
		CLKPR  __reg
		SMCR   __reg
		GPIOR2 __reg
		GPIOR1 __reg
		GPIOR0 __reg
		PLLCSR __reg
		PRR    __reg
	}{
		SPMCSR: 0x57, // Store Program Memory Control Register
		SREG:   0x5f, // Status Register
		SPL:    0x5d, // Stack Pointer
		SPH:    0x5d, // Stack Pointer
		MCUCR:  0x55, // MCU Control Register
		MCUSR:  0x54, // MCU Status Register
		OSCCAL: 0x66, // Oscillator Calibration Value
		CLKPR:  0x61,
		SMCR:   0x53, // Sleep Mode Control Register
		GPIOR2: 0x3a, // General Purpose IO Register 2
		GPIOR1: 0x39, // General Purpose IO Register 1
		GPIOR0: 0x3e, // General Purpose IO Register 0
		PLLCSR: 0x49, // PLL Control And Status Register
		PRR:    0x64, // Power Reduction Register
	}

	// Timer/Counter, 8-bit
	TC8 = struct {
		TIMSK0 __reg
		TIFR0  __reg
		TCCR0A __reg
		TCCR0B __reg
		TCNT0  __reg
		OCR0A  __reg
		OCR0B  __reg
	}{
		TIMSK0: 0x6e, // Timer/Counter0 Interrupt Mask Register
		TIFR0:  0x35, // Timer/Counter0 Interrupt Flag register
		TCCR0A: 0x44, // Timer/Counter  Control Register A
		TCCR0B: 0x45, // Timer/Counter Control Register B
		TCNT0:  0x46, // Timer/Counter0
		OCR0A:  0x47, // Timer/Counter0 Output Compare Register
		OCR0B:  0x48, // Timer/Counter0 Output Compare Register
	}

	// Timer/Counter, 16-bit
	TC16 = struct {
		TIMSK1 __reg
		TIFR1  __reg
		TCCR1A __reg
		TCCR1B __reg
		TCCR1C __reg
		TCNT1L __reg
		TCNT1H __reg
		OCR1AL __reg
		OCR1AH __reg
		OCR1BL __reg
		OCR1BH __reg
		ICR1L  __reg
		ICR1H  __reg
	}{
		TIMSK1: 0x6f, // Timer/Counter Interrupt Mask Register
		TIFR1:  0x36, // Timer/Counter Interrupt Flag register
		TCCR1A: 0x80, // Timer/Counter1 Control Register A
		TCCR1B: 0x81, // Timer/Counter1 Control Register B
		TCCR1C: 0x82, // Timer/Counter1 Control Register C
		TCNT1L: 0x84, // Timer/Counter1 Bytes
		TCNT1H: 0x84, // Timer/Counter1 Bytes
		OCR1AL: 0x88, // Timer/Counter1 Output Compare Register Bytes
		OCR1AH: 0x88, // Timer/Counter1 Output Compare Register Bytes
		OCR1BL: 0x8a, // Timer/Counter1 Output Compare Register Bytes
		OCR1BH: 0x8a, // Timer/Counter1 Output Compare Register Bytes
		ICR1L:  0x86, // Timer/Counter1 Input Capture Register Bytes
		ICR1H:  0x86, // Timer/Counter1 Input Capture Register Bytes
	}

	// Analog-to-Digital Converter
	ADC = struct {
		ADMUX   __reg
		ADCSRA  __reg
		ADCL    __reg
		ADCH    __reg
		ADCSRB  __reg
		DIDR0   __reg
		DIDR1   __reg
		AMP0CSR __reg
		AMP1CSR __reg
		AMP2CSR __reg
	}{
		ADMUX:   0x7c, // The ADC multiplexer Selection Register
		ADCSRA:  0x7a, // The ADC Control and Status register
		ADCL:    0x78, // ADC Data Register Bytes
		ADCH:    0x78, // ADC Data Register Bytes
		ADCSRB:  0x7b, // ADC Control and Status Register B
		DIDR0:   0x7e, // Digital Input Disable Register 0
		DIDR1:   0x7f, // Digital Input Disable Register 0
		AMP0CSR: 0x75,
		AMP1CSR: 0x76,
		AMP2CSR: 0x77,
	}

	// Local Interconnect Network
	LINUART = struct {
		LINCR   __reg
		LINSIR  __reg
		LINENIR __reg
		LINERR  __reg
		LINBTR  __reg
		LINBRRL __reg
		LINBRRH __reg
		LINDLR  __reg
		LINIDR  __reg
		LINSEL  __reg
		LINDAT  __reg
	}{
		LINCR:   0xc8, // LIN Control Register
		LINSIR:  0xc9, // LIN Status and Interrupt Register
		LINENIR: 0xca, // LIN Enable Interrupt Register
		LINERR:  0xcb, // LIN Error Register
		LINBTR:  0xcc, // LIN Bit Timing Register
		LINBRRL: 0xcd, // LIN Baud Rate Register
		LINBRRH: 0xcd, // LIN Baud Rate Register
		LINDLR:  0xcf, // LIN Data Length Register
		LINIDR:  0xd0, // LIN Identifier Register
		LINSEL:  0xd1, // LIN Data Buffer Selection Register
		LINDAT:  0xd2, // LIN Data Register
	}

	// Serial Peripheral Interface
	SPI = struct {
		SPCR __reg
		SPSR __reg
		SPDR __reg
	}{
		SPCR: 0x4c, // SPI Control Register
		SPSR: 0x4d, // SPI Status Register
		SPDR: 0x4e, // SPI Data Register
	}

	// Watchdog Timer
	WDT = struct {
		WDTCSR __reg
	}{
		WDTCSR: 0x60, // Watchdog Timer Control Register
	}

	// External Interrupts
	EXINT = struct {
		EICRA  __reg
		EIMSK  __reg
		EIFR   __reg
		PCICR  __reg
		PCMSK3 __reg
		PCMSK2 __reg
		PCMSK1 __reg
		PCMSK0 __reg
		PCIFR  __reg
	}{
		EICRA:  0x69, // External Interrupt Control Register
		EIMSK:  0x3d, // External Interrupt Mask Register
		EIFR:   0x3c, // External Interrupt Flag Register
		PCICR:  0x68, // Pin Change Interrupt Control Register
		PCMSK3: 0x6d, // Pin Change Mask Register 3
		PCMSK2: 0x6c, // Pin Change Mask Register 2
		PCMSK1: 0x6b, // Pin Change Mask Register 1
		PCMSK0: 0x6a, // Pin Change Mask Register 0
		PCIFR:  0x3b, // Pin Change Interrupt Flag Register
	}

	// EEPROM
	EEPROM = struct {
		EEARL __reg
		EEARH __reg
		EEDR  __reg
		EECR  __reg
	}{
		EEARL: 0x41, // EEPROM Read/Write Access
		EEARH: 0x41, // EEPROM Read/Write Access
		EEDR:  0x40, // EEPROM Data Register
		EECR:  0x3f, // EEPROM Control Register
	}
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_PSCRB    = 0x20 // PSC Reset Behavior
	EXTENDED_PSCRVA   = 0x10 // PSCOUTnA Reset Value
	EXTENDED_PSCRVB   = 0x8  // PSC0UTnB Reset Value
	EXTENDED_BODLEVEL = 0x7  // Brown-out Detector Trigger Level

	// HIGH
	HIGH_RSTDISBL = 0x80 // Reset Disabled (Enable PC6 as i/o pin)
	HIGH_DWEN     = 0x40 // Debug Wire enable
	HIGH_SPIEN    = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON    = 0x10 // Watch-dog Timer always on
	HIGH_EESAVE   = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ   = 0x6  // Select Boot Size
	HIGH_BOOTRST  = 0x1  // Select Reset Vector

	// LOW
	LOW_CKDIV8    = 0x80 // Divide clock by 8 internally
	LOW_CKOUT     = 0x40 // Clock output on PORTD1
	LOW_SUT_CKSEL = 0x3f // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB   = 0x3  // Memory Lock
	LOCKBIT_BLB0 = 0xc  // Boot Loader Protection Mode
	LOCKBIT_BLB1 = 0x30 // Boot Loader Protection Mode
)

// Bitfields for CAN: Controller Area Network
const (
	// CANGCON: CAN General Control Register
	CANGCON_ABRQ   = 0x80 // Abort Request
	CANGCON_OVRQ   = 0x40 // Overload Frame Request
	CANGCON_TTC    = 0x20 // Time Trigger Communication
	CANGCON_SYNTTC = 0x10 // Synchronization of TTC
	CANGCON_LISTEN = 0x8  // Listening Mode
	CANGCON_TEST   = 0x4  // Test Mode
	CANGCON_ENASTB = 0x2  // Enable / Standby
	CANGCON_SWRES  = 0x1  // Software Reset Request

	// CANGSTA: CAN General Status Register
	CANGSTA_OVFG  = 0x40 // Overload Frame Flag
	CANGSTA_TXBSY = 0x10 // Transmitter Busy
	CANGSTA_RXBSY = 0x8  // Receiver Busy
	CANGSTA_ENFG  = 0x4  // Enable Flag
	CANGSTA_BOFF  = 0x2  // Bus Off Mode
	CANGSTA_ERRP  = 0x1  // Error Passive Mode

	// CANGIT: CAN General Interrupt Register Flags
	CANGIT_CANIT  = 0x80 // General Interrupt Flag
	CANGIT_BOFFIT = 0x40 // Bus Off Interrupt Flag
	CANGIT_OVRTIM = 0x20 // Overrun CAN Timer Flag
	CANGIT_BXOK   = 0x10 // Burst Receive Interrupt Flag
	CANGIT_SERG   = 0x8  // Stuff Error General Flag
	CANGIT_CERG   = 0x4  // CRC Error General Flag
	CANGIT_FERG   = 0x2  // Form Error General Flag
	CANGIT_AERG   = 0x1  // Ackknowledgement Error General Flag

	// CANGIE: CAN General Interrupt Enable Register
	CANGIE_ENIT   = 0x80 // Enable all Interrupts
	CANGIE_ENBOFF = 0x40 // Enable Bus Off Interrupt
	CANGIE_ENRX   = 0x20 // Enable Receive Interrupt
	CANGIE_ENTX   = 0x10 // Enable Transmitt Interrupt
	CANGIE_ENERR  = 0x8  // Enable MOb Error Interrupt
	CANGIE_ENBX   = 0x4  // Enable Burst Receive Interrupt
	CANGIE_ENERG  = 0x2  // Enable General Error Interrupt
	CANGIE_ENOVRT = 0x1  // Enable CAN Timer Overrun Interrupt

	// CANEN2: Enable MOb Register 2
	CANEN2_ENMOB = 0x3f // Enable MObs

	// CANIE2: Enable Interrupt MOb Register 2
	CANIE2_IEMOB = 0x3f // Interrupt Enable  MObs

	// CANSIT2: CAN Status Interrupt MOb Register 2
	CANSIT2_SIT = 0x3f // Status of Interrupt MObs

	// CANBT1: CAN Bit Timing Register 1
	CANBT1_BRP = 0x7e // Baud Rate Prescaler bits

	// CANBT2: CAN Bit Timing Register 2
	CANBT2_SJW = 0x60 // Re-Sync Jump Width bits
	CANBT2_PRS = 0xe  // Propagation Time Segment bits

	// CANBT3: CAN Bit Timing Register 3
	CANBT3_PHS2 = 0x70 // Phase Segment 2 bits
	CANBT3_PHS1 = 0xe  // Phase Segment 1 bits
	CANBT3_SMP  = 0x1  // Sample Type

	// CANTCON: Timer Control Register
	CANTCON_TPRSC = 0xff // Timer Control bits

	// CANTIML: Timer Register

	// CANTIMH: Timer Register
	CANTIM_CANTIM = 0xffff // CAN Timer bits

	// CANTTCL: TTC Timer Register

	// CANTTCH: TTC Timer Register
	CANTTC_TIMTTC = 0xffff // TTC Timer Count

	// CANTEC: Transmit Error Counter Register
	CANTEC_TEC = 0xff // Transmit Error Counter bits

	// CANREC: Receive Error Counter Register
	CANREC_REC = 0xff // Receive Error Counter bits

	// CANHPMOB: Highest Priority MOb Register
	CANHPMOB_HPMOB = 0xf0 // Highest Priority MOb Number bits
	CANHPMOB_CGP   = 0xf  // CAN General Purpose bits

	// CANPAGE: Page MOb Register
	CANPAGE_MOBNB = 0xf0 // MOb Number bits
	CANPAGE_AINC  = 0x8  // MOb Data Buffer Auto Increment (Active Low)
	CANPAGE_INDX  = 0x7  // Data Buffer Index bits

	// CANSTMOB: MOb Status Register
	CANSTMOB_DLCW = 0x80 // Data Length Code Warning on MOb
	CANSTMOB_TXOK = 0x40 // Transmit OK on MOb
	CANSTMOB_RXOK = 0x20 // Receive OK on MOb
	CANSTMOB_BERR = 0x10 // Bit Error on MOb
	CANSTMOB_SERR = 0x8  // Stuff Error on MOb
	CANSTMOB_CERR = 0x4  // CRC Error on MOb
	CANSTMOB_FERR = 0x2  // Form Error on MOb
	CANSTMOB_AERR = 0x1  // Ackknowledgement Error on MOb

	// CANCDMOB: MOb Control and DLC Register
	CANCDMOB_CONMOB = 0xc0 // MOb Config bits
	CANCDMOB_RPLV   = 0x20 // Reply Valid
	CANCDMOB_IDE    = 0x10 // Identifier Extension
	CANCDMOB_DLC    = 0xf  // Data Length Code bits

	// CANIDT4: Identifier Tag Register 4
	CANIDT4_IDT    = 0xf8 // Identifier Tag
	CANIDT4_RTRTAG = 0x4  // Remote Transmission Request Tag
	CANIDT4_RB1TAG = 0x2  // Reserved Bit 1 Tag
	CANIDT4_RB0TAG = 0x1  // Reserved Bit 0 Tag

	// CANIDT3: Identifier Tag Register 3
	CANIDT3_IDT = 0xff // Identifier Tag

	// CANIDT2: Identifier Tag Register 2
	CANIDT2_IDT = 0xff // Identifier Tag

	// CANIDT1: Identifier Tag Register 1
	CANIDT1_IDT = 0xff // Identifier Tag

	// CANIDM4: Identifier Mask Register 4
	CANIDM4_IDEMSK = 0x1  // Identifier Extension Mask
	CANIDM4_RTRMSK = 0x4  // Remote Transmission Request Mask
	CANIDM4_IDMSK  = 0xf8 // Identifier Mask

	// CANIDM3: Identifier Mask Register 3
	CANIDM3_IDMSK = 0xff // Identifier Mask

	// CANIDM2: Identifier Mask Register 2
	CANIDM2_IDMSK = 0xff // Identifier Mask

	// CANIDM1: Identifier Mask Register 1
	CANIDM1_IDMSK = 0xff // Identifier Mask

	// CANSTML: Time Stamp Register

	// CANSTMH: Time Stamp Register
	CANSTM_TIMSTM = 0xffff // TIMSTM

	// CANMSG: Message Data Register
	CANMSG_MSG = 0xff // Message Data bits
)

// Bitfields for AC: Analog Comparator
const (
	// AC0CON: Analog Comparator 0 Control Register
	AC0CON_AC0EN   = 0x80 // Analog Comparator 0 Enable Bit
	AC0CON_AC0IE   = 0x40 // Analog Comparator 0 Interrupt Enable Bit
	AC0CON_AC0IS   = 0x30 // Analog Comparator 0  Interrupt Select Bits
	AC0CON_ACCKSEL = 0x8  // Analog Comparator Clock Select
	AC0CON_AC0M    = 0x7  // Analog Comparator 0 Multiplexer Register

	// AC1CON: Analog Comparator 1 Control Register
	AC1CON_AC1EN  = 0x80 // Analog Comparator 1 Enable Bit
	AC1CON_AC1IE  = 0x40 // Analog Comparator 1 Interrupt Enable Bit
	AC1CON_AC1IS  = 0x30 // Analog Comparator 1  Interrupt Select Bit
	AC1CON_AC1ICE = 0x8  // Analog Comparator 1 Interrupt Capture Enable Bit
	AC1CON_AC1M   = 0x7  // Analog Comparator 1 Multiplexer Register

	// AC2CON: Analog Comparator 2 Control Register
	AC2CON_AC2EN = 0x80 // Analog Comparator 2 Enable Bit
	AC2CON_AC2IE = 0x40 // Analog Comparator 2 Interrupt Enable Bit
	AC2CON_AC2IS = 0x30 // Analog Comparator 2  Interrupt Select Bit
	AC2CON_AC2M  = 0x7  // Analog Comparator 2 Multiplexer Register

	// AC3CON: Analog Comparator 3 Control Register
	AC3CON_AC3EN = 0x80 // Analog Comparator 3 Enable Bit
	AC3CON_AC3IE = 0x40 // Analog Comparator 3 Interrupt Enable Bit
	AC3CON_AC3IS = 0x30 // Analog Comparator 3  Interrupt Select Bit
	AC3CON_AC3M  = 0x7  // Analog Comparator 3 Multiplexer Register

	// ACSR: Analog Comparator Status Register
	ACSR_AC3IF = 0x80 // Analog Comparator 3 Interrupt Flag Bit
	ACSR_AC2IF = 0x40 // Analog Comparator 2 Interrupt Flag Bit
	ACSR_AC1IF = 0x20 // Analog Comparator 1  Interrupt Flag Bit
	ACSR_AC0IF = 0x10 // Analog Comparator 0 Interrupt Flag Bit
	ACSR_AC3O  = 0x8  // Analog Comparator 3 Output Bit
	ACSR_AC2O  = 0x4  // Analog Comparator 2 Output Bit
	ACSR_AC1O  = 0x2  // Analog Comparator 1 Output Bit
	ACSR_AC0O  = 0x1  // Analog Comparator 0 Output Bit
)

// Bitfields for DAC: Digital-to-Analog Converter
const (
	// DACL: DAC Data Register

	// DACH: DAC Data Register
	DAC_DACH = 0xff00 // DAC Data Register High Byte Bits
	DAC_DACL = 0xff   // DAC Data Register Low Byte Bits

	// DACON: DAC Control Register
	DACON_DAATE = 0x80 // DAC Auto Trigger Enable Bit
	DACON_DATS  = 0x70 // DAC Trigger Selection Bits
	DACON_DALA  = 0x4  // DAC Left Adjust
	DACON_DAOE  = 0x2  // DAC Output Enable
	DACON_DAEN  = 0x1  // DAC Enable Bit
)

// Bitfields for CPU: CPU Registers
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCSR_SIGRD  = 0x20 // Signature Row Read
	SPMCSR_RWWSRE = 0x10 // Read While Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable

	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// MCUCR: MCU Control Register
	MCUCR_SPIPS = 0x80 // SPI Pin Select
	MCUCR_PUD   = 0x10 // Pull-up disable
	MCUCR_IVSEL = 0x2  // Interrupt Vector Select
	MCUCR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// MCUSR: MCU Status Register
	MCUSR_WDRF  = 0x8 // Watchdog Reset Flag
	MCUSR_BORF  = 0x4 // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2 // External Reset Flag
	MCUSR_PORF  = 0x1 // Power-on reset flag

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL = 0xff // Oscillator Calibration

	// CLKPR
	CLKPR_CLKPCE = 0x80
	CLKPR_CLKPS  = 0xf

	// SMCR: Sleep Mode Control Register
	SMCR_SM = 0xe // Sleep Mode Select bits
	SMCR_SE = 0x1 // Sleep Enable

	// GPIOR2: General Purpose IO Register 2
	GPIOR2_GPIOR = 0xff // General Purpose IO Register 2 bis

	// GPIOR1: General Purpose IO Register 1
	GPIOR1_GPIOR = 0xff // General Purpose IO Register 1 bis

	// GPIOR0: General Purpose IO Register 0
	GPIOR0_GPIOR07 = 0x80 // General Purpose IO Register 0 bit 7
	GPIOR0_GPIOR06 = 0x40 // General Purpose IO Register 0 bit 6
	GPIOR0_GPIOR05 = 0x20 // General Purpose IO Register 0 bit 5
	GPIOR0_GPIOR04 = 0x10 // General Purpose IO Register 0 bit 4
	GPIOR0_GPIOR03 = 0x8  // General Purpose IO Register 0 bit 3
	GPIOR0_GPIOR02 = 0x4  // General Purpose IO Register 0 bit 2
	GPIOR0_GPIOR01 = 0x2  // General Purpose IO Register 0 bit 1
	GPIOR0_GPIOR00 = 0x1  // General Purpose IO Register 0 bit 0

	// PLLCSR: PLL Control And Status Register
	PLLCSR_PLLF  = 0x4 // PLL Factor
	PLLCSR_PLLE  = 0x2 // PLL Enable
	PLLCSR_PLOCK = 0x1 // PLL Lock Detector

	// PRR: Power Reduction Register
	PRR_PRCAN  = 0x40 // Power Reduction CAN
	PRR_PRPSC  = 0x20 // Power Reduction PSC
	PRR_PRTIM1 = 0x10 // Power Reduction Timer/Counter1
	PRR_PRTIM0 = 0x8  // Power Reduction Timer/Counter0
	PRR_PRSPI  = 0x4  // Power Reduction Serial Peripheral Interface
	PRR_PRLIN  = 0x2  // Power Reduction LIN UART
	PRR_PRADC  = 0x1  // Power Reduction ADC
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TIMSK0: Timer/Counter0 Interrupt Mask Register
	TIMSK0_OCIE0B = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag register
	TIFR0_OCF0B = 0x4 // Timer/Counter0 Output Compare Flag 0B
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag 0A
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag

	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_COM0A = 0xc0 // Compare Output Mode, Phase Correct PWM Mode
	TCCR0A_COM0B = 0x30 // Compare Output Mode, Fast PWm
	TCCR0A_WGM0  = 0x3  // Waveform Generation Mode

	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_FOC0A = 0x80 // Force Output Compare A
	TCCR0B_FOC0B = 0x40 // Force Output Compare B
	TCCR0B_WGM02 = 0x8
	TCCR0B_CS0   = 0x7 // Clock Select

	// TCNT0: Timer/Counter0
	TCNT0_TCNT0 = 0xff // Timer/Counter0 bits

	// OCR0A: Timer/Counter0 Output Compare Register
	OCR0A_OCR0A = 0xff // Timer/Counter0 Output Compare bits

	// OCR0B: Timer/Counter0 Output Compare Register
	OCR0B_OCR0B = 0xff // Timer/Counter0 Output Compare bits
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TIMSK1: Timer/Counter Interrupt Mask Register
	TIMSK1_ICIE1  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1B = 0x4  // Timer/Counter1 Output CompareB Match Interrupt Enable
	TIMSK1_OCIE1A = 0x2  // Timer/Counter1 Output CompareA Match Interrupt Enable
	TIMSK1_TOIE1  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter Interrupt Flag register
	TIFR1_ICF1  = 0x20 // Input Capture Flag 1
	TIFR1_OCF1B = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1A = 0x2  // Output Compare Flag 1A
	TIFR1_TOV1  = 0x1  // Timer/Counter1 Overflow Flag

	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_WGM1  = 0x3  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM1  = 0x18 // Waveform Generation Mode
	TCCR1B_CS1   = 0x7  // Prescaler source of Timer/Counter 1

	// TCCR1C: Timer/Counter1 Control Register C
	TCCR1C_FOC1A = 0x80
	TCCR1C_FOC1B = 0x40

	// TCNT1L: Timer/Counter1 Bytes

	// TCNT1H: Timer/Counter1 Bytes
	TCNT1_TCNT1 = 0xffff // Timer/Counter1 bits

	// OCR1AL: Timer/Counter1 Output Compare Register Bytes

	// OCR1AH: Timer/Counter1 Output Compare Register Bytes
	OCR1A_OCR1A = 0xffff // Timer/Counter1 Output Compare bits

	// OCR1BL: Timer/Counter1 Output Compare Register Bytes

	// OCR1BH: Timer/Counter1 Output Compare Register Bytes
	OCR1B_OCR1B = 0xffff // Timer/Counter1 Output Compare bits

	// ICR1L: Timer/Counter1 Input Capture Register Bytes

	// ICR1H: Timer/Counter1 Input Capture Register Bytes
	ICR1_ICR1 = 0xffff // Timer/Counter1 Input Capture bits
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS  = 0xc0 // Reference Selection Bits
	ADMUX_ADLAR = 0x20 // Left Adjust Result
	ADMUX_MUX   = 0x1f // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS  = 0x7  // ADC  Prescaler Select Bits

	// ADCL: ADC Data Register Bytes

	// ADCH: ADC Data Register Bytes
	ADC_ADC = 0x3ff // ADC Data bits

	// ADCSRB: ADC Control and Status Register B
	ADCSRB_ADHSM  = 0x80 // ADC High Speed Mode
	ADCSRB_ISRCEN = 0x40 // Current Source Enable
	ADCSRB_AREFEN = 0x20 // Analog Reference pin Enable
	ADCSRB_ADTS   = 0xf  // ADC Auto Trigger Sources

	// DIDR0: Digital Input Disable Register 0
	DIDR0_ADC7D = 0x80 // ADC7 Digital input Disable
	DIDR0_ADC6D = 0x40 // ADC6 Digital input Disable
	DIDR0_ADC5D = 0x20 // ADC5 Digital input Disable
	DIDR0_ADC4D = 0x10 // ADC4 Digital input Disable
	DIDR0_ADC3D = 0x8  // ADC3 Digital input Disable
	DIDR0_ADC2D = 0x4  // ADC2 Digital input Disable
	DIDR0_ADC1D = 0x2  // ADC1 Digital input Disable
	DIDR0_ADC0D = 0x1  // ADC0 Digital input Disable

	// DIDR1: Digital Input Disable Register 0
	DIDR1_AMP2PD = 0x40 // AMP2P Pin Digital input Disable
	DIDR1_ACMP0D = 0x20 // ACMP0 Pin Digital input Disable
	DIDR1_AMP0PD = 0x10 // AMP0P Pin Digital input Disable
	DIDR1_AMP0ND = 0x8  // AMP0N Pin Digital input Disable
	DIDR1_ADC10D = 0x4  // ADC10 Pin Digital input Disable
	DIDR1_ADC9D  = 0x2  // ADC9 Pin Digital input Disable
	DIDR1_ADC8D  = 0x1  // ADC8 Pin Digital input Disable

	// AMP0CSR
	AMP0CSR_AMP0EN  = 0x80
	AMP0CSR_AMP0IS  = 0x40
	AMP0CSR_AMP0G   = 0x30
	AMP0CSR_AMPCMP0 = 0x8 // Amplifier 0 - Comparator 0 Connection
	AMP0CSR_AMP0TS  = 0x7

	// AMP1CSR
	AMP1CSR_AMP1EN  = 0x80
	AMP1CSR_AMP1IS  = 0x40
	AMP1CSR_AMP1G   = 0x30
	AMP1CSR_AMPCMP1 = 0x8 // Amplifier 1 - Comparator 1 Connection
	AMP1CSR_AMP1TS  = 0x7

	// AMP2CSR
	AMP2CSR_AMP2EN  = 0x80
	AMP2CSR_AMP2IS  = 0x40
	AMP2CSR_AMP2G   = 0x30
	AMP2CSR_AMPCMP2 = 0x8 // Amplifier 2 - Comparator 2 Connection
	AMP2CSR_AMP2TS  = 0x7
)

// Bitfields for LINUART: Local Interconnect Network
const (
	// LINCR: LIN Control Register
	LINCR_LSWRES = 0x80 // Software Reset
	LINCR_LIN13  = 0x40 // LIN Standard
	LINCR_LCONF  = 0x30 // LIN Configuration bits
	LINCR_LENA   = 0x8  // LIN or UART Enable
	LINCR_LCMD   = 0x7  // LIN Command and Mode bits

	// LINSIR: LIN Status and Interrupt Register
	LINSIR_LIDST = 0xe0 // Identifier Status bits
	LINSIR_LBUSY = 0x10 // Busy Signal
	LINSIR_LERR  = 0x8  // Error Interrupt
	LINSIR_LIDOK = 0x4  // Identifier Interrupt
	LINSIR_LTXOK = 0x2  // Transmit Performed Interrupt
	LINSIR_LRXOK = 0x1  // Receive Performed Interrupt

	// LINENIR: LIN Enable Interrupt Register
	LINENIR_LENERR  = 0x8 // Enable Error Interrupt
	LINENIR_LENIDOK = 0x4 // Enable Identifier Interrupt
	LINENIR_LENTXOK = 0x2 // Enable Transmit Performed Interrupt
	LINENIR_LENRXOK = 0x1 // Enable Receive Performed Interrupt

	// LINERR: LIN Error Register
	LINERR_LABORT = 0x80 // Abort Flag
	LINERR_LTOERR = 0x40 // Frame Time Out Error Flag
	LINERR_LOVERR = 0x20 // Overrun Error Flag
	LINERR_LFERR  = 0x10 // Framing Error Flag
	LINERR_LSERR  = 0x8  // Synchronization Error Flag
	LINERR_LPERR  = 0x4  // Parity Error Flag
	LINERR_LCERR  = 0x2  // Checksum Error Flag
	LINERR_LBERR  = 0x1  // Bit Error Flag

	// LINBTR: LIN Bit Timing Register
	LINBTR_LDISR = 0x80 // Disable Bit Timing Resynchronization
	LINBTR_LBT   = 0x3f // LIN Bit Timing bits

	// LINBRRL: LIN Baud Rate Register

	// LINBRRH: LIN Baud Rate Register
	LINBRR_LDIV = 0xfff

	// LINDLR: LIN Data Length Register
	LINDLR_LTXDL = 0xf0 // LIN Transmit Data Length bits
	LINDLR_LRXDL = 0xf  // LIN Receive Data Length bits

	// LINIDR: LIN Identifier Register
	LINIDR_LP  = 0xc0 // Parity bits
	LINIDR_LID = 0x3f // Identifier bit 5 or Data Length bits

	// LINSEL: LIN Data Buffer Selection Register
	LINSEL_LAINC = 0x8 // Auto Increment of Data Buffer Index (Active Low)
	LINSEL_LINDX = 0x7 // FIFO LIN Data Buffer Index bits

	// LINDAT: LIN Data Register
	LINDAT_LDATA = 0xff
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR  = 0x3  // SPI Clock Rate Selects

	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit

	// SPDR: SPI Data Register
	SPDR_SPDR = 0xff // SPI Data bits
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control Register
	WDTCSR_WDIF = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDP  = 0x27 // Watchdog Timer Prescaler Bits
	WDTCSR_WDCE = 0x10 // Watchdog Change Enable
	WDTCSR_WDE  = 0x8  // Watch Dog Enable
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register
	EICRA_ISC3 = 0xc0 // External Interrupt Sense Control Bit
	EICRA_ISC2 = 0x30 // External Interrupt Sense Control Bit
	EICRA_ISC1 = 0xc  // External Interrupt Sense Control 1 Bits
	EICRA_ISC0 = 0x3  // External Interrupt Sense Control 0 Bits

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT = 0xf // External Interrupt Request 3 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF = 0xf // External Interrupt Flags

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE = 0xf // Pin Change Interrupt Enables

	// PCMSK3: Pin Change Mask Register 3
	PCMSK3_PCINT = 0x7 // Pin Change Enable Masks

	// PCMSK2: Pin Change Mask Register 2
	PCMSK2_PCINT = 0xff // Pin Change Enable Masks

	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT = 0xff // Pin Change Enable Masks

	// PCMSK0: Pin Change Mask Register 0
	PCMSK0_PCINT = 0xff // Pin Change Enable Masks

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF = 0xf // Pin Change Interrupt Flags
)

// Bitfields for EEPROM: EEPROM
const (
	// EEARL: EEPROM Read/Write Access

	// EEARH: EEPROM Read/Write Access
	EEAR_EEAR = 0x3ff // EEPROM Address bits

	// EEDR: EEPROM Data Register
	EEDR_EEDR = 0xff // EEPROM Data bits

	// EECR: EEPROM Control Register
	EECR_EEPM  = 0x30 // EEPROM Programming mode
	EECR_EERIE = 0x8  // EEPROM Ready Interrupt Enable
	EECR_EEMWE = 0x4  // EEPROM Master Write Enable
	EECR_EEWE  = 0x2  // EEPROM Write Enable
	EECR_EERE  = 0x1  // EEPROM Read Enable
)
