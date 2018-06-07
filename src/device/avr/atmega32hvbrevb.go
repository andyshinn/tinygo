// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega32HVBrevB.atdf, see http://packs.download.atmel.com/

// +build avr,atmega32hvbrevb

// Device information for the ATmega32HVBrevB.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega32HVBrevB"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET         = 0  // External Pin, Power-on Reset, Brown-out Reset and Watchdog Reset
	IRQ_BPINT         = 1  // Battery Protection Interrupt
	IRQ_VREGMON       = 2  // Voltage regulator monitor interrupt
	IRQ_INT0          = 3  // External Interrupt Request 0
	IRQ_INT1          = 4  // External Interrupt Request 1
	IRQ_INT2          = 5  // External Interrupt Request 2
	IRQ_INT3          = 6  // External Interrupt Request 3
	IRQ_PCINT0        = 7  // Pin Change Interrupt 0
	IRQ_PCINT1        = 8  // Pin Change Interrupt 1
	IRQ_WDT           = 9  // Watchdog Timeout Interrupt
	IRQ_BGSCD         = 10 // Bandgap Buffer Short Circuit Detected
	IRQ_CHDET         = 11 // Charger Detect
	IRQ_TIMER1_IC     = 12 // Timer 1 Input capture
	IRQ_TIMER1_COMPA  = 13 // Timer 1 Compare Match A
	IRQ_TIMER1_COMPB  = 14 // Timer 1 Compare Match B
	IRQ_TIMER1_OVF    = 15 // Timer 1 overflow
	IRQ_TIMER0_IC     = 16 // Timer 0 Input Capture
	IRQ_TIMER0_COMPA  = 17 // Timer 0 Comapre Match A
	IRQ_TIMER0_COMPB  = 18 // Timer 0 Compare Match B
	IRQ_TIMER0_OVF    = 19 // Timer 0 Overflow
	IRQ_TWIBUSCD      = 20 // Two-Wire Bus Connect/Disconnect
	IRQ_TWI           = 21 // Two-Wire Serial Interface
	IRQ_SPI_STC       = 22 // SPI Serial transfer complete
	IRQ_VADC          = 23 // Voltage ADC Conversion Complete
	IRQ_CCADC_CONV    = 24 // Coulomb Counter ADC Conversion Complete
	IRQ_CCADC_REG_CUR = 25 // Coloumb Counter ADC Regular Current
	IRQ_CCADC_ACC     = 26 // Coloumb Counter ADC Accumulator
	IRQ_EE_READY      = 27 // EEPROM Ready
	IRQ_SPM           = 28 // SPM Ready
	IRQ_max           = 28 // Highest interrupt number on this device.
)

// Peripherals
var (
	// Fuses
	FUSE = struct {
		LOW  __reg
		HIGH __reg
	}{
		LOW:  0x0,
		HIGH: 0x1,
	}

	// Lockbits
	LOCKBIT = struct {
		LOCKBIT __reg
	}{
		LOCKBIT: 0x0,
	}

	// Analog-to-Digital Converter
	ADC = struct {
		VADMUX __reg
		VADCL  __reg
		VADCH  __reg
		VADCSR __reg
	}{
		VADMUX: 0x7c, // The VADC multiplexer Selection Register
		VADCL:  0x78, // VADC Data Register  Bytes
		VADCH:  0x78, // VADC Data Register  Bytes
		VADCSR: 0x7a, // The VADC Control and Status register
	}

	// Watchdog Timer
	WDT = struct {
		WDTCSR __reg
	}{
		WDTCSR: 0x60, // Watchdog Timer Control Register
	}

	// FET Control
	FET = struct {
		FCSR __reg
	}{
		FCSR: 0xf0, // FET Control and Status Register
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

	// Coulomb Counter
	COULOMB_COUNTER = struct {
		CADCSRA __reg
		CADCSRB __reg
		CADCSRC __reg
		CADICL  __reg
		CADICH  __reg
		CADAC3  __reg
		CADAC2  __reg
		CADAC1  __reg
		CADAC0  __reg
		CADRCC  __reg
		CADRDC  __reg
	}{
		CADCSRA: 0xe6, // CC-ADC Control and Status Register A
		CADCSRB: 0xe7, // CC-ADC Control and Status Register B
		CADCSRC: 0xe8, // CC-ADC Control and Status Register C
		CADICL:  0xe4, // CC-ADC Instantaneous Current
		CADICH:  0xe4, // CC-ADC Instantaneous Current
		CADAC3:  0xe3, // ADC Accumulate Current
		CADAC2:  0xe2, // ADC Accumulate Current
		CADAC1:  0xe1, // ADC Accumulate Current
		CADAC0:  0xe0, // ADC Accumulate Current
		CADRCC:  0xe9, // CC-ADC Regular Charge Current
		CADRDC:  0xea, // CC-ADC Regular Discharge Current
	}

	// Two Wire Serial Interface
	TWI = struct {
		TWBCSR __reg
		TWAMR  __reg
		TWBR   __reg
		TWCR   __reg
		TWSR   __reg
		TWDR   __reg
		TWAR   __reg
	}{
		TWBCSR: 0xbe, // TWI Bus Control and Status Register
		TWAMR:  0xbd, // TWI (Slave) Address Mask Register
		TWBR:   0xb8, // TWI Bit Rate register
		TWCR:   0xbc, // TWI Control Register
		TWSR:   0xb9, // TWI Status Register
		TWDR:   0xbb, // TWI Data register
		TWAR:   0xba, // TWI (Slave) Address register
	}

	// External Interrupts
	EXINT = struct {
		EICRA  __reg
		EIMSK  __reg
		EIFR   __reg
		PCICR  __reg
		PCIFR  __reg
		PCMSK1 __reg
		PCMSK0 __reg
	}{
		EICRA:  0x69, // External Interrupt Control Register
		EIMSK:  0x3d, // External Interrupt Mask Register
		EIFR:   0x3c, // External Interrupt Flag Register
		PCICR:  0x68, // Pin Change Interrupt Control Register
		PCIFR:  0x3b, // Pin Change Interrupt Flag Register
		PCMSK1: 0x6c, // Pin Change Enable Mask Register 1
		PCMSK0: 0x6b, // Pin Change Enable Mask Register 0
	}

	// Timer/Counter, 16-bit
	TC16 = struct {
		TCCR1B __reg
		TCCR1A __reg
		TCNT1L __reg
		TCNT1H __reg
		OCR1A  __reg
		OCR1B  __reg
		TIMSK1 __reg
		TIFR1  __reg
		TCCR0B __reg
		TCCR0A __reg
		TCNT0L __reg
		TCNT0H __reg
		OCR0A  __reg
		OCR0B  __reg
		TIMSK0 __reg
		TIFR0  __reg
	}{
		TCCR1B: 0x81, // Timer/Counter1 Control Register B
		TCCR1A: 0x80, // Timer/Counter 1 Control Register A
		TCNT1L: 0x84, // Timer Counter 1  Bytes
		TCNT1H: 0x84, // Timer Counter 1  Bytes
		OCR1A:  0x88, // Output Compare Register 1A
		OCR1B:  0x89, // Output Compare Register B
		TIMSK1: 0x6f, // Timer/Counter Interrupt Mask Register
		TIFR1:  0x36, // Timer/Counter Interrupt Flag register
		TCCR0B: 0x45, // Timer/Counter0 Control Register B
		TCCR0A: 0x44, // Timer/Counter 0 Control Register A
		TCNT0L: 0x46, // Timer Counter 0  Bytes
		TCNT0H: 0x46, // Timer Counter 0  Bytes
		OCR0A:  0x48, // Output Compare Register 0A
		OCR0B:  0x49, // Output Compare Register B
		TIMSK0: 0x6e, // Timer/Counter Interrupt Mask Register
		TIFR0:  0x35, // Timer/Counter Interrupt Flag register
	}

	// Cell Balancing
	CELL_BALANCING = struct {
		CBCR __reg
	}{
		CBCR: 0xf1, // Cell Balancing Control Register
	}

	// Battery Protection
	BATTERY_PROTECTION = struct {
		BPPLR  __reg
		BPCR   __reg
		BPHCTR __reg
		BPOCTR __reg
		BPSCTR __reg
		BPCHCD __reg
		BPDHCD __reg
		BPCOCD __reg
		BPDOCD __reg
		BPSCD  __reg
		BPIFR  __reg
		BPIMSK __reg
	}{
		BPPLR:  0xfe, // Battery Protection Parameter Lock Register
		BPCR:   0xfd, // Battery Protection Control Register
		BPHCTR: 0xfc, // Battery Protection Short-current Timing Register
		BPOCTR: 0xfb, // Battery Protection Over-current Timing Register
		BPSCTR: 0xfa, // Battery Protection Short-current Timing Register
		BPCHCD: 0xf9, // Battery Protection Charge-High-current Detection Level Register
		BPDHCD: 0xf8, // Battery Protection Discharge-High-current Detection Level Register
		BPCOCD: 0xf7, // Battery Protection Charge-Over-current Detection Level Register
		BPDOCD: 0xf6, // Battery Protection Discharge-Over-current Detection Level Register
		BPSCD:  0xf5, // Battery Protection Short-Circuit Detection Level Register
		BPIFR:  0xf3, // Battery Protection Interrupt Flag Register
		BPIMSK: 0xf2, // Battery Protection Interrupt Mask Register
	}

	// Charger Detect
	CHARGER_DETECT = struct {
		CHGDCSR __reg
	}{
		CHGDCSR: 0xd4, // Charger Detect Control and Status Register
	}

	// Voltage Regulator
	VOLTAGE_REGULATOR = struct {
		ROCR __reg
	}{
		ROCR: 0xc8, // Regulator Operating Condition Register
	}

	// Bandgap
	BANDGAP = struct {
		BGCSR __reg
		BGCRR __reg
		BGCCR __reg
	}{
		BGCSR: 0xd2, // Bandgap Control and Status Register
		BGCRR: 0xd1, // Bandgap Calibration of Resistor Ladder
		BGCCR: 0xd0, // Bandgap Calibration Register
	}

	// CPU Registers
	CPU = struct {
		SREG    __reg
		SPL     __reg
		SPH     __reg
		MCUCR   __reg
		MCUSR   __reg
		FOSCCAL __reg
		OSICSR  __reg
		SMCR    __reg
		GPIOR2  __reg
		GPIOR1  __reg
		GPIOR0  __reg
		DIDR0   __reg
		PRR0    __reg
		CLKPR   __reg
	}{
		SREG:    0x5f, // Status Register
		SPL:     0x5d, // Stack Pointer
		SPH:     0x5d, // Stack Pointer
		MCUCR:   0x55, // MCU Control Register
		MCUSR:   0x54, // MCU Status Register
		FOSCCAL: 0x66, // Fast Oscillator Calibration Value
		OSICSR:  0x37, // Oscillator Sampling Interface Control and Status Register
		SMCR:    0x53, // Sleep Mode Control Register
		GPIOR2:  0x4b, // General Purpose IO Register 2
		GPIOR1:  0x4a, // General Purpose IO Register 1
		GPIOR0:  0x3e, // General Purpose IO Register 0
		DIDR0:   0x7e, // Digital Input Disable Register
		PRR0:    0x64, // Power Reduction Register 0
		CLKPR:   0x61, // Clock Prescale Register
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
		PINC  __reg
	}{
		PORTA: 0x22, // Port A Data Register
		DDRA:  0x21, // Port A Data Direction Register
		PINA:  0x20, // Port A Input Pins
		PORTB: 0x25, // Port B Data Register
		DDRB:  0x24, // Port B Data Direction Register
		PINB:  0x23, // Port B Input Pins
		PORTC: 0x28, // Port C Data Register
		PINC:  0x26, // Port C Input Pins
	}

	// Bootloader
	BOOT_LOAD = struct {
		SPMCSR __reg
	}{
		SPMCSR: 0x57, // Store Program Memory Control and Status Register
	}
)

// Bitfields for FUSE: Fuses
const (
	// LOW
	LOW_WDTON  = 0x80 // Watch-dog Timer always on
	LOW_EESAVE = 0x40 // Preserve EEPROM through the Chip Erase cycle
	LOW_SPIEN  = 0x20 // Serial program downloading (SPI) enabled
	LOW_SUT    = 0x1c // Select start-up time
	LOW_OSCSEL = 0x3  // Oscillator select

	// HIGH
	HIGH_DUVRDINIT = 0x10 // DUVR mode on
	HIGH_DWEN      = 0x8  // Debug Wire enable
	HIGH_BOOTSZ    = 0x6  // Select Boot Size
	HIGH_BOOTRST   = 0x1  // Boot Reset vector Enabled
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB   = 0x3  // Memory Lock
	LOCKBIT_BLB0 = 0xc  // Boot Loader Protection Mode
	LOCKBIT_BLB1 = 0x30 // Boot Loader Protection Mode
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// VADMUX: The VADC multiplexer Selection Register
	VADMUX_VADMUX = 0xf // Analog Channel and Gain Selection Bits

	// VADCSR: The VADC Control and Status register
	VADCSR_VADEN   = 0x8 // VADC Enable
	VADCSR_VADSC   = 0x4 // VADC Satrt Conversion
	VADCSR_VADCCIF = 0x2 // VADC Conversion Complete Interrupt Flag
	VADCSR_VADCCIE = 0x1 // VADC Conversion Complete Interrupt Enable
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

// Bitfields for FET: FET Control
const (
	// FCSR: FET Control and Status Register
	FCSR_DUVRD = 0x8 // Deep Under-Voltage Recovery Disable
	FCSR_CPS   = 0x4 // Current Protection Status
	FCSR_DFE   = 0x2 // Discharge FET Enable
	FCSR_CFE   = 0x1 // Charge FET Enable
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

// Bitfields for EEPROM: EEPROM
const (
	// EECR: EEPROM Control Register
	EECR_EEPM  = 0x30
	EECR_EERIE = 0x8 // EEProm Ready Interrupt Enable
	EECR_EEMPE = 0x4 // EEPROM Master Write Enable
	EECR_EEPE  = 0x2 // EEPROM Write Enable
	EECR_EERE  = 0x1 // EEPROM Read Enable
)

// Bitfields for COULOMB_COUNTER: Coulomb Counter
const (
	// CADCSRA: CC-ADC Control and Status Register A
	CADCSRA_CADEN  = 0x80 // When the CADEN bit is cleared (zero), the CC-ADC is disabled. When the CADEN bit is set (one), the CC-ADC will continuously measure the voltage drop over the external sense resistor RSENSE. In Power-down, only the Regular Current detection is active. In Power-off, the CC-ADC is always disabled.
	CADCSRA_CADPOL = 0x40
	CADCSRA_CADUB  = 0x20 // CC_ADC Update Busy
	CADCSRA_CADAS  = 0x18 // CC_ADC Accumulate Current Select Bits
	CADCSRA_CADSI  = 0x6  // The CADSI bits determine the current sampling interval for the Regular Current detection in Power-down mode. The actual settings remain to be determined.
	CADCSRA_CADSE  = 0x1  // When the CADSE bit is written to one, the ongoing CC-ADC conversion is aborted, and the CC-ADC enters Regular Current detection mode.

	// CADCSRB: CC-ADC Control and Status Register B
	CADCSRB_CADACIE = 0x40
	CADCSRB_CADRCIE = 0x20 // Regular Current Interrupt Enable
	CADCSRB_CADICIE = 0x10 // CAD Instantenous Current Interrupt Enable
	CADCSRB_CADACIF = 0x4  // CC-ADC Accumulate Current Interrupt Flag
	CADCSRB_CADRCIF = 0x2  // CC-ADC Accumulate Current Interrupt Flag
	CADCSRB_CADICIF = 0x1  // CC-ADC Instantaneous Current Interrupt Flag

	// CADCSRC: CC-ADC Control and Status Register C
	CADCSRC_CADVSE = 0x1 // CC-ADC Voltage Scaling Enable
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWBCSR: TWI Bus Control and Status Register
	TWBCSR_TWBCIF = 0x80 // TWI Bus Connect/Disconnect Interrupt Flag
	TWBCSR_TWBCIE = 0x40 // TWI Bus Connect/Disconnect Interrupt Enable
	TWBCSR_TWBDT  = 0x6  // TWI Bus Disconnect Time-out Period
	TWBCSR_TWBCIP = 0x1  // TWI Bus Connect/Disconnect Interrupt Polarity

	// TWAMR: TWI (Slave) Address Mask Register
	TWAMR_TWAM = 0xfe

	// TWCR: TWI Control Register
	TWCR_TWINT = 0x80 // TWI Interrupt Flag
	TWCR_TWEA  = 0x40 // TWI Enable Acknowledge Bit
	TWCR_TWSTA = 0x20 // TWI Start Condition Bit
	TWCR_TWSTO = 0x10 // TWI Stop Condition Bit
	TWCR_TWWC  = 0x8  // TWI Write Collition Flag
	TWCR_TWEN  = 0x4  // TWI Enable Bit
	TWCR_TWIE  = 0x1  // TWI Interrupt Enable

	// TWSR: TWI Status Register
	TWSR_TWS  = 0xf8 // TWI Status
	TWSR_TWPS = 0x3  // TWI Prescaler

	// TWAR: TWI (Slave) Address register
	TWAR_TWA   = 0xfe // TWI (Slave) Address register Bits
	TWAR_TWGCE = 0x1  // TWI General Call Recognition Enable Bit
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register
	EICRA_ISC3 = 0xc0 // External Interrupt Sense Control 3 Bits
	EICRA_ISC2 = 0x30 // External Interrupt Sense Control 2 Bits
	EICRA_ISC1 = 0xc  // External Interrupt Sense Control 1 Bits
	EICRA_ISC0 = 0x3  // External Interrupt Sense Control 0 Bits

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT = 0xf // External Interrupt Request 3 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF = 0xf // External Interrupt Flags

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE = 0x3 // Pin Change Interrupt Enables

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF = 0x3 // Pin Change Interrupt Flags
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_CS = 0x7 // Clock Select1 bis

	// TCCR1A: Timer/Counter 1 Control Register A
	TCCR1A_TCW1  = 0x80 // Timer/Counter Width
	TCCR1A_ICEN1 = 0x40 // Input Capture Mode Enable
	TCCR1A_ICNC1 = 0x20 // Input Capture Noise Canceler
	TCCR1A_ICES1 = 0x10 // Input Capture Edge Select
	TCCR1A_ICS1  = 0x8  // Input Capture Select
	TCCR1A_WGM10 = 0x1  // Waveform Generation Mode

	// TIMSK1: Timer/Counter Interrupt Mask Register
	TIMSK1_ICIE1  = 0x8 // Timer/Counter n Input Capture Interrupt Enable
	TIMSK1_OCIE1B = 0x4 // Timer/Counter1 Output Compare B Interrupt Enable
	TIMSK1_OCIE1A = 0x2 // Timer/Counter1 Output Compare A Interrupt Enable
	TIMSK1_TOIE1  = 0x1 // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter Interrupt Flag register
	TIFR1_ICF1  = 0x8 // Timer/Counter 1 Input Capture Flag
	TIFR1_OCF1B = 0x4 // Timer/Counter1 Output Compare Flag B
	TIFR1_OCF1A = 0x2 // Timer/Counter1 Output Compare Flag A
	TIFR1_TOV1  = 0x1 // Timer/Counter1 Overflow Flag

	// TCCR0B: Timer/Counter0 Control Register B
	TCCR0B_CS02 = 0x4 // Clock Select0 bit 2
	TCCR0B_CS01 = 0x2 // Clock Select0 bit 1
	TCCR0B_CS00 = 0x1 // Clock Select0 bit 0

	// TCCR0A: Timer/Counter 0 Control Register A
	TCCR0A_TCW0  = 0x80 // Timer/Counter Width
	TCCR0A_ICEN0 = 0x40 // Input Capture Mode Enable
	TCCR0A_ICNC0 = 0x20 // Input Capture Noise Canceler
	TCCR0A_ICES0 = 0x10 // Input Capture Edge Select
	TCCR0A_ICS0  = 0x8  // Input Capture Select
	TCCR0A_WGM00 = 0x1  // Waveform Generation Mode

	// TIMSK0: Timer/Counter Interrupt Mask Register
	TIMSK0_ICIE0  = 0x8 // Timer/Counter n Input Capture Interrupt Enable
	TIMSK0_OCIE0B = 0x4 // Timer/Counter0 Output Compare B Interrupt Enable
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare A Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter Interrupt Flag register
	TIFR0_ICF0  = 0x8 // Timer/Counter 0 Input Capture Flag
	TIFR0_OCF0B = 0x4 // Timer/Counter0 Output Compare Flag B
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag A
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag
)

// Bitfields for CELL_BALANCING: Cell Balancing
const (
	// CBCR: Cell Balancing Control Register
	CBCR_CBE = 0xf // Cell Balancing Enables
)

// Bitfields for BATTERY_PROTECTION: Battery Protection
const (
	// BPPLR: Battery Protection Parameter Lock Register
	BPPLR_BPPLE = 0x2 // Battery Protection Parameter Lock Enable
	BPPLR_BPPL  = 0x1 // Battery Protection Parameter Lock

	// BPCR: Battery Protection Control Register
	BPCR_EPID = 0x20 // External Protection Input Disable
	BPCR_SCD  = 0x10 // Short Circuit Protection Disabled
	BPCR_DOCD = 0x8  // Discharge Over-current Protection Disabled
	BPCR_COCD = 0x4  // Charge Over-current Protection Disabled
	BPCR_DHCD = 0x2  // Discharge High-current Protection Disable
	BPCR_CHCD = 0x1  // Charge High-current Protection Disable

	// BPIFR: Battery Protection Interrupt Flag Register
	BPIFR_SCIF  = 0x10 // Short-circuit Protection Activated Interrupt Flag
	BPIFR_DOCIF = 0x8  // Discharge Over-current Protection Activated Interrupt Flag
	BPIFR_COCIF = 0x4  // Charge Over-current Protection Activated Interrupt Flag
	BPIFR_DHCIF = 0x2  // Disharge High-current Protection Activated Interrupt
	BPIFR_CHCIF = 0x1  // Charge High-current Protection Activated Interrupt

	// BPIMSK: Battery Protection Interrupt Mask Register
	BPIMSK_SCIE  = 0x10 // Short-circuit Protection Activated Interrupt Enable
	BPIMSK_DOCIE = 0x8  // Discharge Over-current Protection Activated Interrupt Enable
	BPIMSK_COCIE = 0x4  // Charge Over-current Protection Activated Interrupt Enable
	BPIMSK_DHCIE = 0x2  // Discharger High-current Protection Activated Interrupt
	BPIMSK_CHCIE = 0x1  // Charger High-current Protection Activated Interrupt
)

// Bitfields for CHARGER_DETECT: Charger Detect
const (
	// CHGDCSR: Charger Detect Control and Status Register
	CHGDCSR_BATTPVL = 0x10 // BATT Pin Voltage Level
	CHGDCSR_CHGDISC = 0xc  // Charger Detect Interrupt Sense Control
	CHGDCSR_CHGDIF  = 0x2  // Charger Detect Interrupt Flag
	CHGDCSR_CHGDIE  = 0x1  // Charger Detect Interrupt Enable
)

// Bitfields for VOLTAGE_REGULATOR: Voltage Regulator
const (
	// ROCR: Regulator Operating Condition Register
	ROCR_ROCS   = 0x80 // ROC Status
	ROCR_ROCD   = 0x10 // ROC Disable
	ROCR_ROCWIF = 0x2  // ROC Warning Interrupt Flag
	ROCR_ROCWIE = 0x1  // ROC Warning Interrupt Enable
)

// Bitfields for BANDGAP: Bandgap
const (
	// BGCSR: Bandgap Control and Status Register
	BGCSR_BGD     = 0x20 // Bandgap Disable
	BGCSR_BGSCDE  = 0x10 // Bandgap Short Circuit Detection Enabled
	BGCSR_BGSCDIF = 0x2  // Bandgap Short Circuit Detection Interrupt Flag
	BGCSR_BGSCDIE = 0x1  // Bandgap Short Circuit Detection Interrupt Enable

	// BGCCR: Bandgap Calibration Register
	BGCCR_BGCC = 0x3f // BG Calibration of PTAT Current Bits
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

	// MCUCR: MCU Control Register
	MCUCR_CKOE  = 0x20 // Clock Output Enable
	MCUCR_PUD   = 0x10 // Pull-up disable
	MCUCR_IVSEL = 0x2  // Interrupt Vector Select
	MCUCR_IVCE  = 0x1  // Interrupt Vector Change Enable

	// MCUSR: MCU Status Register
	MCUSR_OCDRF = 0x10 // OCD Reset Flag
	MCUSR_WDRF  = 0x8  // Watchdog Reset Flag
	MCUSR_BODRF = 0x4  // Brown-out Reset Flag
	MCUSR_EXTRF = 0x2  // External Reset Flag
	MCUSR_PORF  = 0x1  // Power-on reset flag

	// OSICSR: Oscillator Sampling Interface Control and Status Register
	OSICSR_OSISEL0 = 0x10 // Oscillator Sampling Interface Select 0
	OSICSR_OSIST   = 0x2  // Oscillator Sampling Interface Status
	OSICSR_OSIEN   = 0x1  // Oscillator Sampling Interface Enable

	// SMCR: Sleep Mode Control Register
	SMCR_SM = 0xe // Sleep Mode Select bits
	SMCR_SE = 0x1 // Sleep Enable

	// DIDR0: Digital Input Disable Register
	DIDR0_PA1DID = 0x2 // When this bit is written logic one, the digital input buffer of the corresponding V_ADC pin is disabled.
	DIDR0_PA0DID = 0x1 // When this bit is written logic one, the digital input buffer of the corresponding V_ADC pin is disabled.

	// PRR0: Power Reduction Register 0
	PRR0_PRTWI  = 0x40 // Power Reduction TWI
	PRR0_PRVRM  = 0x20 // Power Reduction Voltage Regulator Monitor
	PRR0_PRSPI  = 0x8  // Power reduction SPI
	PRR0_PRTIM1 = 0x4  // Power Reduction Timer/Counter1
	PRR0_PRTIM0 = 0x2  // Power Reduction Timer/Counter0
	PRR0_PRVADC = 0x1  // Power Reduction V-ADC

	// CLKPR: Clock Prescale Register
	CLKPR_CLKPCE = 0x80 // Clock Prescaler Change Enable
	CLKPR_CLKPS  = 0x3  // Clock Prescaler Select Bits
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control and Status Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read-While-Write Section Busy
	SPMCSR_SIGRD  = 0x20 // Signature Row Read
	SPMCSR_RWWSRE = 0x10 // Read-While-Write Section Read Enable
	SPMCSR_LBSET  = 0x8  // Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
)
