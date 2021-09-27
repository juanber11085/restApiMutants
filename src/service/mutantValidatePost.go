package service

import (
	"fmt"
	"main/src/repository"
	"main/src/repository/entity"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type requestMutantValidate struct {
	Dna []string `json:"dna"`
}

var reSequence = regexp.MustCompile(`(A){4,}|(T){4,}|(C){4,}|(G){4,}`)
var isValidateData bool

func init() {
	isValidateData = true
}

func MutantValidatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		dnas := requestMutantValidate{}
		c.Bind(&dnas)
		flagIsMutant := isMutant(dnas.Dna)
		if !isValidateData {
			c.JSON(400, gin.H{
				"message": "Bad request.",
			})
		} else {
			var mutantSave = entity.Mutants{
				Id:       strings.Join(dnas.Dna, ","),
				IsMutant: flagIsMutant,
			}
			var flagSave = saveAdn(mutantSave)
			if !flagSave {
				c.JSON(500, gin.H{
					"message": "Internal Server Error.",
				})
			} else {
				if flagIsMutant {
					c.JSON(200, gin.H{
						"message": "All Ok.",
					})
				} else {
					c.JSON(403, gin.H{
						"message": "Forbidden.",
					})
				}
			}
		}
	}
}

func isMutant(dna []string) bool {
	var mutant = getMutantsById(strings.Join(dna, ","))
	if (entity.Mutants{}) != mutant {
		return mutant.IsMutant
	}
	reTypeCharacters := regexp.MustCompile(`\b[ATCG]{2,}\b`)
	var dataSequence int = 0
	var arrayHorizontal []string
	var arrayOblique []map[string]string
	for i := 0; i < len(dna); i++ {
		dna[i] = strings.ToUpper(dna[i])
		var element = dna[i]
		if !reTypeCharacters.MatchString(element) {
			isValidateData = false
			break
		} else {
			dataSequence += len(reSequence.FindAllString(element, -1))
			if dataSequence > 1 {
				break
			}
		}
		var elementArray = strings.Split(element, "")
		for j := 0; j < len(elementArray); j++ {
			var itemElementArray = elementArray[j]
			var itemMap map[string]string = make(map[string]string)
			var posHorizontal int = i
			var posVertical int = j
			if len(arrayHorizontal) <= j {
				arrayHorizontal = append(arrayHorizontal, elementArray[j])
			} else {
				var elementEdit = arrayHorizontal[j]
				elementEdit += itemElementArray
				arrayHorizontal[j] = elementEdit
			}
			if i == 0 {
				if (len(elementArray) - j) > 3 {
					itemMap[fmt.Sprintf("%d-%d-r", (posHorizontal+1), (posVertical+1))] = itemElementArray
					arrayOblique = append(arrayOblique, itemMap)
				} else {
					itemMap[fmt.Sprintf("%d-%d-l", (posHorizontal+1), (posVertical-1))] = itemElementArray
					arrayOblique = append(arrayOblique, itemMap)
				}
			} else {
				if j == 0 {
					if (len(dna) - i) > 3 {
						itemMap[fmt.Sprintf("%d-%d-r", (posHorizontal+1), (posVertical+1))] = itemElementArray
						arrayOblique = append(arrayOblique, itemMap)
					} else {
						arrayOblique = AddItemsArrayOblique(i, j, itemElementArray, arrayOblique)
					}
				} else if j == (len(elementArray) - 1) {
					if (len(dna) - i) > 3 {
						itemMap[fmt.Sprintf("%d-%d-l", (posHorizontal+1), (posVertical-1))] = itemElementArray
						arrayOblique = append(arrayOblique, itemMap)
					} else {
						arrayOblique = AddItemsArrayOblique(i, j, itemElementArray, arrayOblique)
					}
				} else {
					arrayOblique = AddItemsArrayOblique(i, j, itemElementArray, arrayOblique)
				}
			}
		}
	}
	if !isValidateData {
		return false
	} else {
		if dataSequence > 1 {
			return true
		} else {
			if dataSequence = validateArrayHorizontal(arrayHorizontal, dataSequence); dataSequence > 1 {
				return true
			} else {
				if dataSequence = validateArrayOblique(arrayOblique, dataSequence); dataSequence > 1 {
					return true
				} else {
					return false
				}
			}
		}

	}
}

func AddItemsArrayOblique(posHorizontal int, posVertical int, valueItem string, arrayOblique []map[string]string) []map[string]string {
	for i := 0; i < len(arrayOblique); i++ {
		var itemArrayOblique = arrayOblique[i]
		for key, value := range itemArrayOblique {
			var splitPositionTable = strings.Split(key, "-")
			positionHorizontalItem, _ := strconv.Atoi(splitPositionTable[0])
			positionVerticalItem, _ := strconv.Atoi(splitPositionTable[1])
			if posHorizontal == positionHorizontalItem && posVertical == positionVerticalItem {
				if splitPositionTable[2] == "r" {
					itemArrayOblique[fmt.Sprintf("%d-%d-r", (positionHorizontalItem+1), (positionVerticalItem+1))] = value + valueItem
					delete(itemArrayOblique, key)
				} else if splitPositionTable[2] == "l" {
					if positionVerticalItem == 0 {
						itemArrayOblique[fmt.Sprintf("%d-%d-e", (positionHorizontalItem+1), (positionVerticalItem-1))] = value + valueItem
					} else {
						itemArrayOblique[fmt.Sprintf("%d-%d-l", (positionHorizontalItem+1), (positionVerticalItem-1))] = value + valueItem
					}
					delete(itemArrayOblique, key)
				}
			}
		}
	}
	return arrayOblique
}

func validateArrayHorizontal(arrayHorizontal []string, dataSequence int) int {
	for _, itemHorizontal := range arrayHorizontal {
		dataSequence += len(reSequence.FindAllString(itemHorizontal, -1))
		if dataSequence > 1 {
			return dataSequence
		}
	}
	return dataSequence
}

func validateArrayOblique(arrayOblique []map[string]string, dataSequence int) int {
	for _, itemOblique := range arrayOblique {
		for _, value := range itemOblique {
			dataSequence += len(reSequence.FindAllString(value, -1))
			if dataSequence > 1 {
				return dataSequence
			}
		}
	}
	return dataSequence
}

func getMutantsById(dnaId string) entity.Mutants {
	var mutant, err = repository.GetItem(dnaId)
	if err != nil {
		return entity.Mutants{}
	} else {
		return mutant
	}
}

func saveAdn(mutant entity.Mutants) bool {
	var err = repository.PutItem(mutant)
	if err != nil {
		return false
	} else {
		return true
	}
}