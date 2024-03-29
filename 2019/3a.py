import re

w1path = ['R1009', 'U263', 'L517', 'U449', 'L805', 'D78', 'L798', 'D883', 'L777', 'D562', 'R652', 'D348', 'R999', 'D767', 'L959', 'U493', 'R59', 'D994', 'L225', 'D226', 'R634', 'D200', 'R953', 'U343', 'L388', 'U158', 'R943', 'U544', 'L809', 'D785', 'R618', 'U499', 'L476', 'U600', 'L452', 'D693', 'L696', 'U764', 'L927', 'D346', 'L863', 'D458', 'L789', 'U268', 'R586', 'U884', 'L658', 'D371', 'L910', 'U178', 'R524', 'U169', 'R973', 'D326', 'R483', 'U233', 'R26', 'U807', 'L246', 'D711', 'L641', 'D75', 'R756', 'U365', 'R203', 'D377', 'R624', 'U430', 'L422', 'U367', 'R547', 'U294', 'L916', 'D757', 'R509', 'D332', 'R106', 'D401', 'L181', 'U5', 'L443', 'U197', 'R406', 'D829', 'R878', 'U35', 'L958', 'U31', 'L28', 'D362', 'R188', 'D582', 'R358', 'U750', 'R939', 'D491', 'R929', 'D513', 'L541', 'U418', 'R861', 'D639', 'L917', 'U582', 'R211', 'U725', 'R711', 'D718', 'L673', 'U921', 'L157', 'U83', 'L199', 'U501', 'L66', 'D993', 'L599', 'D947', 'L26', 'U237', 'L981', 'U833', 'L121', 'U25', 'R641', 'D372', 'L757', 'D645', 'R287', 'U390', 'R274', 'U964', 'R288', 'D209', 'R109', 'D364', 'R983', 'U715', 'L315', 'U758', 'R36', 'D500', 'R626', 'U893', 'L840', 'U716', 'L606', 'U831', 'L969', 'D643', 'L300', 'D838', 'R31', 'D751', 'L632', 'D702', 'R468', 'D7', 'L169', 'U149', 'R893', 'D33', 'R816', 'D558', 'R152', 'U489', 'L237', 'U415', 'R434', 'D472', 'L198', 'D874', 'L351', 'U148', 'R761', 'U809', 'R21', 'D25', 'R586', 'D338', 'L568', 'U20', 'L157', 'U221', 'L26', 'U424', 'R261', 'D227', 'L551', 'D754', 'L90', 'U110', 'L791', 'U433', 'R840', 'U323', 'R240', 'U124', 'L723', 'D418', 'R938', 'D173', 'L160', 'U293', 'R773', 'U204', 'R192', 'U958', 'L472', 'D703', 'R556', 'D168', 'L263', 'U574', 'L845', 'D932', 'R165', 'D348', 'R811', 'D834', 'R960', 'U877', 'R935', 'D141', 'R696', 'U748', 'L316', 'U236', 'L796', 'D566', 'R524', 'U449', 'R378', 'U480', 'L79', 'U227', 'R867', 'D185', 'R474', 'D757', 'R366', 'U153', 'R882', 'U252', 'R861', 'U900', 'R28', 'U381', 'L845', 'U642', 'L849', 'U352', 'R134', 'D294', 'R788', 'D406', 'L693', 'D697', 'L433', 'D872', 'R78', 'D364', 'R240', 'U995', 'R48', 'D681', 'R727', 'D825', 'L583', 'U44', 'R743', 'D929', 'L616', 'D262', 'R997', 'D15', 'R575', 'U341', 'R595', 'U889', 'R254', 'U76', 'R962', 'D944', 'R724', 'D261', 'R608', 'U753', 'L389', 'D324', 'L569', 'U308', 'L488', 'D358', 'L695', 'D863', 'L712', 'D978', 'R149', 'D177', 'R92']
w2path = ['L1003', 'D960', 'L10', 'D57', 'R294', 'U538', 'R867', 'D426', 'L524', 'D441', 'R775', 'U308', 'R577', 'D785', 'R495', 'U847', 'R643', 'D895', 'R448', 'U685', 'L253', 'U312', 'L312', 'U753', 'L89', 'U276', 'R799', 'D923', 'L33', 'U595', 'R400', 'U111', 'L664', 'D542', 'R171', 'U709', 'L809', 'D713', 'L483', 'U918', 'L14', 'U854', 'L150', 'D69', 'L158', 'D500', 'L91', 'D800', 'R431', 'D851', 'L798', 'U515', 'L107', 'U413', 'L94', 'U390', 'L17', 'U221', 'L999', 'D546', 'L191', 'U472', 'L568', 'U114', 'L913', 'D743', 'L713', 'D215', 'L569', 'D674', 'L869', 'U549', 'L789', 'U259', 'L330', 'D76', 'R243', 'D592', 'L646', 'U880', 'L363', 'U542', 'L464', 'D955', 'L107', 'U473', 'R818', 'D786', 'R852', 'U968', 'R526', 'D78', 'L275', 'U891', 'R480', 'U991', 'L981', 'D391', 'R83', 'U691', 'R689', 'D230', 'L217', 'D458', 'R10', 'U736', 'L317', 'D145', 'R902', 'D428', 'R344', 'U334', 'R131', 'D739', 'R438', 'D376', 'L652', 'U304', 'L332', 'D452', 'R241', 'D783', 'R82', 'D317', 'R796', 'U323', 'R287', 'D487', 'L302', 'D110', 'R233', 'U631', 'R584', 'U973', 'L878', 'D834', 'L930', 'U472', 'R120', 'U78', 'R806', 'D21', 'L521', 'U988', 'R251', 'D817', 'R44', 'D789', 'R204', 'D669', 'R616', 'D96', 'R624', 'D891', 'L532', 'U154', 'R438', 'U469', 'R785', 'D431', 'R945', 'U649', 'R670', 'D11', 'R840', 'D521', 'L235', 'D69', 'L551', 'D266', 'L454', 'U807', 'L885', 'U590', 'L647', 'U763', 'R449', 'U194', 'R68', 'U809', 'L884', 'U962', 'L476', 'D648', 'L139', 'U96', 'L300', 'U351', 'L456', 'D202', 'R168', 'D698', 'R161', 'U834', 'L273', 'U47', 'L8', 'D157', 'L893', 'D200', 'L454', 'U723', 'R886', 'U92', 'R474', 'U262', 'L190', 'U110', 'L407', 'D723', 'R786', 'D786', 'L572', 'D915', 'L904', 'U744', 'L820', 'D663', 'R205', 'U878', 'R186', 'U247', 'L616', 'D386', 'R582', 'U688', 'L349', 'D399', 'R702', 'U132', 'L276', 'U866', 'R851', 'D633', 'R468', 'D263', 'R678', 'D96', 'L50', 'U946', 'R349', 'D482', 'R487', 'U525', 'R464', 'U977', 'L499', 'D187', 'R546', 'U708', 'L627', 'D470', 'R673', 'D886', 'L375', 'U616', 'L503', 'U38', 'L775', 'D8', 'L982', 'D556', 'R159', 'U680', 'L124', 'U777', 'L640', 'D607', 'R248', 'D671', 'L65', 'D290', 'R445', 'U778', 'L650', 'U679', 'L846', 'D1', 'L769', 'U659', 'R734', 'D962', 'R588', 'U178', 'R888', 'D753', 'R223', 'U318', 'L695', 'D586', 'R430', 'D61', 'R105', 'U801', 'R953', 'U721', 'L856', 'U769', 'R937', 'D335', 'R895']
w1test = ['R75','D30','R83','U83','L12','D49','R71','U7','L72']
w2test = ['U62','R66','U55','R34','D71','R55','D58','R83']

def calcWireTrail(path):
  trail = [[0, 0]]
  for move in path:
    direction = re.findall('[A-Z]+', move)[0]
    length = int(re.findall('[0-9]+', move)[0])
    increaseX = direction == 'R'
    decreaseX = direction == 'L'
    increaseY = direction == 'U'
    decreaseY = direction == 'D'
    currentPos = trail[len(trail) - 1]

    lineTrail = [currentPos]
    if increaseX:
      for step in range(length):
        newPos = [currentPos[0] + step + 1, currentPos[1]]
        lineTrail.append(newPos)
    elif decreaseX:
      for step in range(length):
        newPos = [currentPos[0] - step + 1, currentPos[1]]
        lineTrail.append(newPos)
    elif increaseY:
      for step in range(length):
        newPos = [currentPos[0], currentPos[1] + step + 1]
        lineTrail.append(newPos)
    elif decreaseY:
      for step in range(length):
        newPos = [currentPos[0], currentPos[1] - step + 1]
        lineTrail.append(newPos)
    trail += lineTrail
  return trail

def findIntersections(wire1, wire2):
  intersections = []
  skip = False
  for wire1Position in wire1:
    if skip or (wire1Position[0] == 0 and wire1Position[1] == 0):
      skip = False
      continue

    for wire2Position in wire2:
      print(f'checking if {wire1Position} == {wire2Position}')

      if wire1Position[0] == wire2Position[0] and wire1Position[1] == wire2Position[1]:
        print(f'intersection found {wire1Position} == {wire2Position}')
        intersections.append(wire1Position)
        skip = True
        break
  return intersections

def getManhattanDistance(intersection):
  return abs(0 - intersection[0]) + abs(0 - intersection[1])

w1trail = calcWireTrail(w1test)
w2trail = calcWireTrail(w2test)

distances = []
for intersection in findIntersections(w1trail, w2trail):
  distances.append(getManhattanDistance(intersection))
print(distances)

# while True:
#   line = input().strip().split(',')
#   print(line)
