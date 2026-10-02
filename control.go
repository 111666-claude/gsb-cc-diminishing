// Package control 是控制递减：按类别分别计数、窗口内减半、免疫窗口与查询代价。
// 缺陷：递减用倒数而不是减半、类别不区分、窗口不重置、免疫不生效、
// 每次施加扫描全部单位。
package control

// Tracker 是控制递减台账。
type Tracker struct {
	windowMs int64
	immuneMs int64
	count    map[string]int
	last     map[string]int64
	scanned  int
}

// New 建台账：windowMs 是递减窗口，immuneMs 是免疫时长。
func New(windowMs, immuneMs int64) *Tracker {
	return &Tracker{windowMs: windowMs, immuneMs: immuneMs,
		count: map[string]int{}, last: map[string]int64{}}
}

// Apply 施加一次控制，返回实际生效时长。
// 缺陷：只按单位累计次数（不分类别）、用 duration/(count+1) 递减、
// 窗口从不重置、没有免疫。
func (t *Tracker) Apply(unitID, category string, durationMs, nowMs int64) int64 {
	t.scanned += len(t.count)
	_ = category
	t.last[unitID] = nowMs
	count := t.count[unitID]
	t.count[unitID] = count + 1
	return durationMs / int64(count+1)
}

// Count 返回某个单位某个类别当前的递减计数。
func (t *Tracker) Count(unitID, category string) int {
	_ = category
	return t.count[unitID]
}

// Window 是递减窗口长度。
func (t *Tracker) Window() int64 { return t.windowMs }

// Immune 是免疫时长。
func (t *Tracker) Immune() int64 { return t.immuneMs }

// Scanned 是累计扫描量（规模观测）。
func (t *Tracker) Scanned() int { return t.scanned }
