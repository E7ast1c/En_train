package domain

type IrregularVerbs struct {
	Id                       int    `db:"id"`
	Infinitive               string `db:"infinitive"`
	InfinitiveTranscript     string `db:"infinitive_transcript"`
	PastTense                string `db:"past_tense"`
	PastTenseTranscript      string `db:"past_tense_transcript"`
	PastParticiple           string `db:"past_participle"`
	PastParticipleTranscript string `db:"past_participle_transcript"`
	Translate                string `db:"translate"`
}
