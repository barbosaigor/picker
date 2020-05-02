# Picker
Picker implements an algorithm that randomly selects an element, taking into account each element probability. 
Higher weight elements are more likely to be choosen.

## Documentation

**Add** attaches an element. 

| Element | Weight |
| ------------- | -------------:|
| A | 80 |
| B | 90 |
| C | 10 |


```golang
p := picker.New()
p.Add("a", 80)
p.Add("b", 90)
p.Add("c", 10)
```

**Roll Dice** selects a random element taking into account each element probability. 
Higher weight elements are more likely to be choosen.

```golang
p := picker.New()
p.Add("a", 80)
p.Add("b", 90)
p.Add("c", 10)
element, err := p.RollDice()
```

**Reset** removes all elements and resets picker variables. 

```golang
p.Reset()
```