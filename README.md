# TSIA2 Practice Pacing: A Checkpoint Budget

This tiny Go package accompanies an educational method for pacing a practice set without turning every question into a stopwatch event. It divides a self-selected practice budget into cumulative checkpoints so the learner can review pace at a few planned moments.

It is not an official testing-time specification. Learners should follow the rules and interface shown by their testing program. The method here is for independent practice and error review.

## Why checkpoints help

Checking the clock after every item consumes attention and can create unnecessary urgency. Ignoring time completely can allow one difficult item to absorb an entire practice block. Checkpoints provide a middle path:

1. Choose a question count and a practice-time budget.
2. Divide both into three or four blocks.
3. Compare actual progress with the cumulative target only at each checkpoint.
4. Record the cause of any large difference after the set.

For a 24-question practice set in 36 minutes, three equal checkpoints are:

| Checkpoint | Cumulative questions | Cumulative minutes |
|---|---:|---:|
| 1 | 8 | 12 |
| 2 | 16 | 24 |
| 3 | 24 | 36 |

These are review markers, not per-question limits. A learner can spend more time on one item and less on another while remaining near the block target.

## Use an intervention rule

Decide before the set what will happen when a checkpoint is missed. A simple rule is:

> If I am more than one question behind at a checkpoint, I will mark the next unresolved item, choose the best supported option, and preserve time for the remaining questions.

The rule should be specific enough to act on, but not so rigid that it overrides careful reading.

## Diagnose the reason after the set

“Too slow” is not a useful diagnosis. Tag the actual cause:

- **Translation delay:** converting words into an equation took too long.
- **Evidence loop:** rereading the same lines without forming a claim.
- **Calculation repair:** arithmetic had to be restarted after a sign or unit error.
- **Choice conflict:** two options remained because the deciding constraint was not identified.
- **Attention reset:** progress paused because focus drifted.

The next practice session should target the dominant cause. More speed drills will not fix a translation problem if the learner has not practiced defining variables and units.

## Package example

```go
points := pacing.Build(24, 36, 3)
// [{8 12} {16 24} {24 36}]
```

`Build` returns cumulative integer checkpoints. It returns `nil` for non-positive inputs. When totals do not divide evenly, integer checkpoints stay monotonic and the final checkpoint always equals the requested totals.

## A final review question

After each checkpoint, ask: “Was I behind because I was reasoning carefully, or because I was repeating work without gaining evidence?” The first may be a reasonable tradeoff. The second identifies a process to change.

For additional timed practice, use the publisher’s **[TSI Practice Test app on Google Play](https://play.google.com/store/apps/details?id=com.tsi.practicetest&hl=en_US)**.

Official app listing: https://play.google.com/store/apps/details?id=com.tsi.practicetest&hl=en_US

## Ownership disclosure

This original educational package and guide were published by the app publisher. The Google Play link identifies the related official listing.
