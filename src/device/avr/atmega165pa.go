// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega165PA.atdf, see http://packs.download.atmel.com/

// +build avr,atmega165pa

// Device information for the ATmega165PA.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega165PA"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin,Power-on Reset,Brown-out Reset,Watchdog Reset,and JTAG AVR Reset. See Datasheet.
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_PCINT0       = 2  // Pin Change Interrupt Request 0
	IRQ_PCINT1       = 3  // Pin Change Interrupt Request 1
	IRQ_TIMER2_COMP  = 4  // Timer/Counter2 Compare Match
	IRQ_TIMER2_OVF   = 5  // Timer/Counter2 Overflow
	IRQ_TIMER1_CAPT  = 6  // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 7  // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 8  // Timer/Counter Compare Match B
	IRQ_TIMER1_OVF   = 9  // Timer/Counter1 Overflow
	IRQ_TIMER0_COMP  = 10 // Timer/Counter0 Compare Match
	IRQ_TIMER0_OVF   = 11 // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 12 // SPI Serial Transfer Complete
	IRQ_USART0_RX    = 13 // USART0, Rx Complete
	IRQ_USART0_UDRE  = 14 // USART0 Data register Empty
	IRQ_USART0_TX    = 15 // USART0, Tx Complete
	IRQ_USI_START    = 16 // USI Start Condition
	IRQ_USI_OVERFLOW = 17 // USI Overflow
	IRQ_ANALOG_COMP  = 18 // Analog Comparator
	IRQ_ADC          = 19 // ADC Conversion Complete
	IRQ_EE_READY     = 20 // EEPROM Ready
	IRQ_SPM_READY    = 21 // Store Program Memory Read
	IRQ_max          = 21 // Highest interrupt number on this device.
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

	// Timer/Counter, 8-bit
	TC8 = struct {
		TCCR0A __reg
		TCNT0  __reg
		OCR0A  __reg
		TIMSK0 __reg
		TIFR0  __reg
	}{
		TCCR0A: 0x44, // Timer/Counter0 Control Register
		TCNT0:  0x46, // Timer/Counter0
		OCR0A:  0x47, // Timer/Counter0 Output Compare Register
		TIMSK0: 0x6e, // Timer/Counter0 Interrupt Mask Register
		TIFR0:  0x35, // Timer/Counter0 Interrupt Flag register
	}

	// Timer/Counter, 16-bit
	TC16 = struct {
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
		TIMSK1 __reg
		TIFR1  __reg
	}{
		TCCR1A: 0x80, // Timer/Counter1 Control Register A
		TCCR1B: 0x81, // Timer/Counter1 Control Register B
		TCCR1C: 0x82, // Timer/Counter 1 Control Register C
		TCNT1L: 0x84, // Timer/Counter1  Bytes
		TCNT1H: 0x84, // Timer/Counter1  Bytes
		OCR1AL: 0x88, // Timer/Counter1 Output Compare Register A  Bytes
		OCR1AH: 0x88, // Timer/Counter1 Output Compare Register A  Bytes
		OCR1BL: 0x8a, // Timer/Counter1 Output Compare Register B  Bytes
		OCR1BH: 0x8a, // Timer/Counter1 Output Compare Register B  Bytes
		ICR1L:  0x86, // Timer/Counter1 Input Capture Register  Bytes
		ICR1H:  0x86, // Timer/Counter1 Input Capture Register  Bytes
		TIMSK1: 0x6f, // Timer/Counter1 Interrupt Mask Register
		TIFR1:  0x36, // Timer/Counter1 Interrupt Flag register
	}

	// Timer/Counter, 8-bit Async
	TC8_ASYNC = struct {
		TCCR2A __reg
		TCNT2  __reg
		OCR2A  __reg
		TIMSK2 __reg
		TIFR2  __reg
		ASSR   __reg
	}{
		TCCR2A: 0xb0, // Timer/Counter2 Control Register
		TCNT2:  0xb2, // Timer/Counter2
		OCR2A:  0xb3, // Timer/Counter2 Output Compare Register
		TIMSK2: 0x70, // Timer/Counter2 Interrupt Mask register
		TIFR2:  0x37, // Timer/Counter2 Interrupt Flag Register
		ASSR:   0xb6, // Asynchronous Status Register
	}

	// Watchdog Timer
	WDT = struct {
		WDTCR __reg
	}{
		WDTCR: 0x60, // Watchdog Timer Control Register
	}

	// EEPROM
	EEPROM = struct {
		EEARL __reg
		EEARH __reg
		EEDR  __reg
		EECR  __reg
	}{
		EEARL: 0x41, // EEPROM Address Register  Bytes
		EEARH: 0x41, // EEPROM Address Register  Bytes
		EEDR:  0x40, // EEPROM Data Register
		EECR:  0x3f, // EEPROM Control Register
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

	// I/O Port
	PORT = struct {
		PORTA __reg
		DDRA  __reg
		PINA  __reg
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
		PORTF __reg
		DDRF  __reg
		PINF  __reg
		PORTG __reg
		DDRG  __reg
		PING  __reg
	}{
		PORTA: 0x22, // Port A Data Register
		DDRA:  0x21, // Port A Data Direction Register
		PINA:  0x20, // Port A Input Pins
		PORTB: 0x25, // Port B Data Register
		DDRB:  0x24, // Port B Data Direction Register
		PINB:  0x23, // Port B Input Pins
		PORTC: 0x28, // Port C Data Register
		DDRC:  0x27, // Port C Data Direction Register
		PINC:  0x26, // Port C Input Pins
		PORTD: 0x2b, // Port D Data Register
		DDRD:  0x2a, // Port D Data Direction Register
		PIND:  0x29, // Port D Input Pins
		PORTE: 0x2e, // Data Register, Port E
		DDRE:  0x2d, // Data Direction Register, Port E
		PINE:  0x2c, // Input Pins, Port E
		PORTF: 0x31, // Data Register, Port F
		DDRF:  0x30, // Data Direction Register, Port F
		PINF:  0x2f, // Input Pins, Port F
		PORTG: 0x34, // Port G Data Register
		DDRG:  0x33, // Port G Data Direction Register
		PING:  0x32, // Port G Input Pins
	}

	// Analog Comparator
	AC = struct {
		ACSR  __reg
		DIDR1 __reg
	}{
		ACSR:  0x50, // Analog Comparator Control And Status Register
		DIDR1: 0x7f, // Digital Input Disable Register 1
	}

	// JTAG Interface
	JTAG = struct {
		OCDR __reg
	}{
		OCDR: 0x51, // On-Chip Debug Related Register in I/O Memory
	}

	// Universal Serial Interface
	USI = struct {
		USIDR __reg
		USISR __reg
		USICR __reg
	}{
		USIDR: 0xba, // USI Data Register
		USISR: 0xb9, // USI Status Register
		USICR: 0xb8, // USI Control Register
	}

	// Analog-to-Digital Converter
	ADC = struct {
		ADMUX  __reg
		ADCSRA __reg
		ADCL   __reg
		ADCH   __reg
		DIDR0  __reg
	}{
		ADMUX:  0x7c, // The ADC multiplexer Selection Register
		ADCSRA: 0x7a, // The ADC Control and Status register
		ADCL:   0x78, // ADC Data Register  Bytes
		ADCH:   0x78, // ADC Data Register  Bytes
		DIDR0:  0x7e, // Digital Input Disable Register 0
	}

	// Bootloader
	BOOT_LOAD = struct {
		SPMCSR __reg
	}{
		SPMCSR: 0x57, // Store Program Memory Control Register
	}

	// USART
	USART = struct {
		UDR0   __reg
		UCSR0A __reg
		UCSR0B __reg
		UCSR0C __reg
		UBRR0L __reg
		UBRR0H __reg
	}{
		UDR0:   0xc6, // USART I/O Data Register
		UCSR0A: 0xc0, // USART Control and Status Register A
		UCSR0B: 0xc1, // USART Control and Status Register B
		UCSR0C: 0xc2, // USART Control and Status Register C
		UBRR0L: 0xc4, // USART Baud Rate Register  Bytes
		UBRR0H: 0xc4, // USART Baud Rate Register  Bytes
	}

	// External Interrupts
	EXINT = struct {
		EICRA  __reg
		EIMSK  __reg
		EIFR   __reg
		PCMSK1 __reg
		PCMSK0 __reg
	}{
		EICRA:  0x69, // External Interrupt Control Register
		EIMSK:  0x3d, // External Interrupt Mask Register
		EIFR:   0x3c, // External Interrupt Flag Register
		PCMSK1: 0x6c, // Pin Change Mask Register 1
		PCMSK0: 0x6b, // Pin Change Mask Register 0
	}

	// CPU Registers
	CPU = struct {
		SREG   __reg
		SPL    __reg
		SPH    __reg
		OSCCAL __reg
		CLKPR  __reg
		PRR    __reg
		SMCR   __reg
		GPIOR2 __reg
		GPIOR1 __reg
		GPIOR0 __reg
	}{
		SREG:   0x5f, // Status Register
		SPL:    0x5d, // Stack Pointer
		SPH:    0x5d, // Stack Pointer
		OSCCAL: 0x66, // Oscillator Calibration Value
		CLKPR:  0x61, // Clock Prescale Register
		PRR:    0x64, // Power Reduction Register
		SMCR:   0x53, // Sleep Mode Control Register
		GPIOR2: 0x4b, // General Purpose IO Register 2
		GPIOR1: 0x4a, // General Purpose IO Register 1
		GPIOR0: 0x3e, // General Purpose IO Register 0
	}
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_BODLEVEL = 0xe // Brown-out Detector trigger level
	EXTENDED_RSTDISBL = 0x1 // Disable external reset

	// HIGH
	HIGH_OCDEN   = 0x80 // On-Chip Debug Enabled
	HIGH_JTAGEN  = 0x40 // JTAG Interface Enabled
	HIGH_SPIEN   = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON   = 0x10 // Watchdog timer always on
	HIGH_EESAVE  = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ  = 0x6  // Select Boot Size
	HIGH_BOOTRST = 0x1  // Boot Reset vector Enabled

	// LOW
	LOW_CKDIV8    = 0x80 // Divide clock by 8 internally
	LOW_CKOUT     = 0x40 // Clock output on PORTE7
	LOW_SUT_CKSEL = 0x3f // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB   = 0x3  // Memory Lock
	LOCKBIT_BLB0 = 0xc  // Boot Loader Protection Mode
	LOCKBIT_BLB1 = 0x30 // Boot Loader Protection Mode
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// TCCR0A: Timer/Counter0 Control Register
	TCCR0A_FOC0A = 0x80 // Force Output Compare
	TCCR0A_WGM00 = 0x40 // Waveform Generation Mode 0
	TCCR0A_COM0A = 0x30 // Compare Match Output Modes
	TCCR0A_WGM01 = 0x8  // Waveform Generation Mode 1
	TCCR0A_CS0   = 0x7  // Clock Selects

	// TIMSK0: Timer/Counter0 Interrupt Mask Register
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare Match Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag register
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag 0
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_WGM1  = 0x3  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM1  = 0x18 // Waveform Generation Mode
	TCCR1B_CS1   = 0x7  // Prescaler source of Timer/Counter 1

	// TCCR1C: Timer/Counter 1 Control Register C
	TCCR1C_FOC1A = 0x80 // Force Output Compare 1A
	TCCR1C_FOC1B = 0x40 // Force Output Compare 1B

	// TIMSK1: Timer/Counter1 Interrupt Mask Register
	TIMSK1_ICIE1  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1B = 0x4  // Timer/Counter1 Output Compare B Match Interrupt Enable
	TIMSK1_OCIE1A = 0x2  // Timer/Counter1 Output Compare A Match Interrupt Enable
	TIMSK1_TOIE1  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter1 Interrupt Flag register
	TIFR1_ICF1  = 0x20 // Input Capture Flag 1
	TIFR1_OCF1B = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1A = 0x2  // Output Compare Flag 1A
	TIFR1_TOV1  = 0x1  // Timer/Counter1 Overflow Flag
)

// Bitfields for TC8_ASYNC: Timer/Counter, 8-bit Async
const (
	// TCCR2A: Timer/Counter2 Control Register
	TCCR2A_FOC2A = 0x80 // Force Output Compare A
	TCCR2A_WGM20 = 0x40 // Waveform Generation Mode
	TCCR2A_COM2A = 0x30 // Compare Output Mode bits
	TCCR2A_WGM21 = 0x8  // Waveform Generation Mode
	TCCR2A_CS2   = 0x7  // Clock Select bits

	// TIMSK2: Timer/Counter2 Interrupt Mask register
	TIMSK2_OCIE2A = 0x2 // Timer/Counter2 Output Compare Match Interrupt Enable
	TIMSK2_TOIE2  = 0x1 // Timer/Counter2 Overflow Interrupt Enable

	// TIFR2: Timer/Counter2 Interrupt Flag Register
	TIFR2_OCF2A = 0x2 // Timer/Counter2 Output Compare Flag 2
	TIFR2_TOV2  = 0x1 // Timer/Counter2 Overflow Flag

	// ASSR: Asynchronous Status Register
	ASSR_EXCLK  = 0x10 // Enable External Clock Interrupt
	ASSR_AS2    = 0x8  // AS2: Asynchronous Timer/Counter2
	ASSR_TCN2UB = 0x4  // TCN2UB: Timer/Counter2 Update Busy
	ASSR_OCR2UB = 0x2  // Output Compare Register2 Update Busy
	ASSR_TCR2UB = 0x1  // TCR2UB: Timer/Counter Control Register2 Update Busy
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCR: Watchdog Timer Control Register
	WDTCR_WDCE = 0x10 // Watchdog Change Enable
	WDTCR_WDE  = 0x8  // Watch Dog Enable
	WDTCR_WDP  = 0x7  // Watch Dog Timer Prescaler bits
)

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EERIE = 0x8 // EEPROM Ready Interrupt Enable
	EECR_EEMWE = 0x4 // EEPROM Master Write Enable
	EECR_EEWE  = 0x2 // EEPROM Write Enable
	EECR_EERE  = 0x1 // EEPROM Read Enable
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
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR: Analog Comparator Control And Status Register
	ACSR_ACD  = 0x80 // Analog Comparator Disable
	ACSR_ACBG = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACO  = 0x20 // Analog Compare Output
	ACSR_ACI  = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACIE = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIC = 0x4  // Analog Comparator Input Capture Enable
	ACSR_ACIS = 0x3  // Analog Comparator Interrupt Mode Select bits

	// DIDR1: Digital Input Disable Register 1
	DIDR1_AIN1D = 0x2 // AIN1 Digital Input Disable
	DIDR1_AIN0D = 0x1 // AIN0 Digital Input Disable
)

// Bitfields for USI: Universal Serial Interface
const (
	// USISR: USI Status Register
	USISR_USISIF = 0x80 // Start Condition Interrupt Flag
	USISR_USIOIF = 0x40 // Counter Overflow Interrupt Flag
	USISR_USIPF  = 0x20 // Stop Condition Flag
	USISR_USIDC  = 0x10 // Data Output Collision
	USISR_USICNT = 0xf  // USI Counter Value Bits

	// USICR: USI Control Register
	USICR_USISIE = 0x80 // Start Condition Interrupt Enable
	USICR_USIOIE = 0x40 // Counter Overflow Interrupt Enable
	USICR_USIWM  = 0x30 // USI Wire Mode Bits
	USICR_USICS  = 0xc  // USI Clock Source Select Bits
	USICR_USICLK = 0x2  // Clock Strobe
	USICR_USITC  = 0x1  // Toggle Clock Port Pin
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

	// DIDR0: Digital Input Disable Register 0
	DIDR0_ADC7D = 0x80 // ADC7 Digital input Disable
	DIDR0_ADC6D = 0x40 // ADC6 Digital input Disable
	DIDR0_ADC5D = 0x20 // ADC5 Digital input Disable
	DIDR0_ADC4D = 0x10 // ADC4 Digital input Disable
	DIDR0_ADC3D = 0x8  // ADC3 Digital input Disable
	DIDR0_ADC2D = 0x4  // ADC2 Digital input Disable
	DIDR0_ADC1D = 0x2  // ADC1 Digital input Disable
	DIDR0_ADC0D = 0x1  // ADC0 Digital input Disable
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCSR_RWWSRE = 0x10 // Read While Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
)

// Bitfields for USART: USART
const (
	// UCSR0A: USART Control and Status Register A
	UCSR0A_RXC0  = 0x80 // USART Receive Complete
	UCSR0A_TXC0  = 0x40 // USART Transmit Complete
	UCSR0A_UDRE0 = 0x20 // USART Data Register Empty
	UCSR0A_FE0   = 0x10 // Framing Error
	UCSR0A_DOR0  = 0x8  // Data OverRun
	UCSR0A_UPE0  = 0x4  // USART Parity Error
	UCSR0A_U2X0  = 0x2  // Double the USART Transmission Speed
	UCSR0A_MPCM0 = 0x1  // Multi-processor Communication Mode

	// UCSR0B: USART Control and Status Register B
	UCSR0B_RXCIE0 = 0x80 // RX Complete Interrupt Enable
	UCSR0B_TXCIE0 = 0x40 // TX Complete Interrupt Enable
	UCSR0B_UDRIE0 = 0x20 // USART Data Register Empty Interrupt Enable
	UCSR0B_RXEN0  = 0x10 // Receiver Enable
	UCSR0B_TXEN0  = 0x8  // Transmitter Enable
	UCSR0B_UCSZ02 = 0x4  // Character Size
	UCSR0B_RXB80  = 0x2  // Receive Data Bit 8
	UCSR0B_TXB80  = 0x1  // Transmit Data Bit 8

	// UCSR0C: USART Control and Status Register C
	UCSR0C_UMSEL0 = 0x40 // USART Mode Select
	UCSR0C_UPM0   = 0x30 // Parity Mode Bits
	UCSR0C_USBS0  = 0x8  // Stop Bit Select
	UCSR0C_UCSZ0  = 0x6  // Character Size
	UCSR0C_UCPOL0 = 0x1  // Clock Polarity
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register
	EICRA_ISC01 = 0x2 // External Interrupt Sense Control 0 Bit 1
	EICRA_ISC00 = 0x1 // External Interrupt Sense Control 0 Bit 0

	// EIMSK: External Interrupt Mask Register
	EIMSK_PCIE = 0x30 // Pin Change Interrupt Enables
	EIMSK_INT0 = 0x1  // External Interrupt Request 0 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_PCIF  = 0x30 // Pin Change Interrupt Flags
	EIFR_INTF0 = 0x1  // External Interrupt Flag 0

	// PCMSK1: Pin Change Mask Register 1
	PCMSK1_PCINT = 0xff // Pin Change Enable Masks

	// PCMSK0: Pin Change Mask Register 0
	PCMSK0_PCINT = 0xff // Pin Change Enable Masks
)

// Bitfields for CPU: CPU Registers
const (
	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL = 0xff // Oscillator Calibration

	// CLKPR: Clock Prescale Register
	CLKPR_CLKPCE = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPS  = 0xf  // Clock Prescaler Select Bits

	// PRR: Power Reduction Register
	PRR_PRTIM1   = 0x8 // Power Reduction Timer/Counter1
	PRR_PRSPI    = 0x4 // Power Reduction Serial Peripheral Interface
	PRR_PRUSART0 = 0x2 // Power Reduction USART
	PRR_PRADC    = 0x1 // Power Reduction ADC

	// SMCR: Sleep Mode Control Register
	SMCR_SM = 0xe // Sleep Mode Select bits
	SMCR_SE = 0x1 // Sleep Enable
)
